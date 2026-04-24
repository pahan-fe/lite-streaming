package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pahan-fe/lite-streaming/backend/internal/config"
	"github.com/pahan-fe/lite-streaming/backend/internal/queue"
	"github.com/pahan-fe/lite-streaming/backend/internal/repository"
	"github.com/pahan-fe/lite-streaming/backend/internal/storage"
	"github.com/pahan-fe/lite-streaming/backend/internal/transcoder"
)

func processMessage(body []byte, repo *repository.VideoRepository, str *storage.S3Storage, tc *transcoder.Transcoder) (err error) {
	var task map[string]string

	unmarshalErr := json.Unmarshal(body, &task)
	if unmarshalErr != nil {
		return fmt.Errorf("Failed to parse message: %v", unmarshalErr)
	}

	tmpDir, tmpDirErr := os.MkdirTemp("", "transcode-")
	if tmpDirErr != nil {
		return fmt.Errorf("Failed to create temporary directory: %v", tmpDirErr)
	}
	defer os.RemoveAll(tmpDir)

	videoId := task["video_id"]
	defer func() {
		if err != nil {
			repo.UpdateStatus(videoId, "failed")
		}
	}()

	video, videoErr := repo.GetByID(videoId)
	if videoErr != nil {
		return fmt.Errorf("Failed to get video by ID: %v", videoErr)
	}

	originVideo, originVideoErr := str.Get(video.S3RawKey)
	if originVideoErr != nil {
		return fmt.Errorf("Failed to get origin video: %v", originVideoErr)
	}

	repo.UpdateStatus(video.ID, "processing")

	var file = tmpDir + "/input.mp4"

	os.WriteFile(file, originVideo, 0644)

	var hlsDir = tmpDir + "/hls"
	transcodeErr := tc.TranscodeToHLS(file, hlsDir)
	if transcodeErr != nil {
		return fmt.Errorf("Failed to transcode video: %v", transcodeErr)
	}

	segments, segmentsErr := os.ReadDir(hlsDir)
	if segmentsErr != nil {
		return fmt.Errorf("Failed to read temporary directory: %v", segmentsErr)
	}

	var hlsKey = "videos/" + videoId + "/hls"

	for _, segment := range segments {
		segmentPath := hlsDir + "/" + segment.Name()

		fileContent, readErr := os.ReadFile(segmentPath)
		if readErr != nil {
			return fmt.Errorf("Failed to read segment: %v", readErr)
		}

		var ext string
		if strings.HasSuffix(segment.Name(), ".m3u8") {
			ext = "application/vnd.apple.mpegurl"
		} else {
			ext = "video/mp2t"
		}

		uploadErr := str.Upload(hlsKey+"/"+segment.Name(), fileContent, ext)
		if uploadErr != nil {
			return fmt.Errorf("Failed to upload segment: %v", uploadErr)
		}
	}

	repo.UpdateStatus(video.ID, "ready")

	return nil
}

func main() {
	cfg := config.Load()

	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	str, storageErr := storage.NewS3Storage(&cfg)
	if storageErr != nil {
		log.Fatalf("Failed to initialize storage: %v", storageErr)
	}

	mq, queueErr := queue.NewRabbitMQ(&cfg)
	if queueErr != nil {
		log.Fatalf("Failed to initialize queue: %v", queueErr)
	}

	repo := repository.NewVideoRepository(db)
	tc := transcoder.NewTranscoder()

	msgs, _ := mq.Consume("transcode")
	for msg := range msgs {
		err := processMessage(msg.Body, repo, str, tc)
		if err != nil {
			log.Printf("Failed to process message: %v", err)
		}
	}
}
