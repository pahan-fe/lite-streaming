package main

import (
	"encoding/json"
	"os"
	"strings"
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pahan-fe/lite-streaming/backend/internal/queue"
	"github.com/pahan-fe/lite-streaming/backend/internal/config"
	"github.com/pahan-fe/lite-streaming/backend/internal/storage"
	"github.com/pahan-fe/lite-streaming/backend/internal/transcoder"
	"github.com/pahan-fe/lite-streaming/backend/internal/repository"
)

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
		var task map[string]string

		unmarshalErr := json.Unmarshal(msg.Body, &task)
		if unmarshalErr != nil {
			log.Printf("Failed to parse message: %v", unmarshalErr)
			continue
		}

		tmpDir, tmpDirErr := os.MkdirTemp("", "transcode-")                                                       
		if tmpDirErr != nil {
			log.Printf("Failed to create temporary directory: %v", tmpDirErr)
			continue
		}

		videoId := task["video_id"]

		video, videoErr := repo.GetByID(videoId)
		if videoErr != nil {
			log.Printf("Failed to get video by ID: %v", videoErr)
			continue
		}

		originVideo, originVideoErr := str.Get(video.S3RawKey)
		if originVideoErr != nil {
			log.Printf("Failed to get origin video: %v", originVideoErr)
			continue
		}

		repo.UpdateStatus(video.ID, "processing")

		var file = tmpDir+"/input.mp4"

		os.WriteFile(file, originVideo, 0644)

		var hlsDir = tmpDir+"/hls"
		transcodeErr := tc.TranscodeToHLS(file, hlsDir)
		if transcodeErr != nil {
			log.Printf("Failed to transcode video: %v", transcodeErr)
			continue
		}

		segments, segmentsErr := os.ReadDir(hlsDir)
		if segmentsErr != nil {
			log.Printf("Failed to read temporary directory: %v", segmentsErr)
			continue
		}

		var hlsKey = "videos/" + videoId + "/hls"

		for _, segment := range segments {
			segmentPath := hlsDir + "/" + segment.Name()

            fileContent, readErr := os.ReadFile(segmentPath)
            if readErr != nil {
                log.Printf("Failed to read segment: %v", readErr)
                continue
            }

			var ext string
			if strings.HasSuffix(segment.Name(), ".m3u8") {
				ext = "application/vnd.apple.mpegurl"
			} else {
				ext = "video/mp2t"
			}

			uploadErr := str.Upload(hlsKey+"/"+segment.Name(), fileContent, ext)
			if uploadErr != nil {
				log.Printf("Failed to upload segment: %v", uploadErr)
				continue
			}
		}

		repo.UpdateStatus(video.ID, "ready")

		os.RemoveAll(tmpDir)	
	}
}