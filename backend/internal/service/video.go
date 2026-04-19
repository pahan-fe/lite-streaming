package service

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pahan-fe/lite-streaming/backend/internal/model"
	"github.com/pahan-fe/lite-streaming/backend/internal/queue"
	"github.com/pahan-fe/lite-streaming/backend/internal/repository"
	"github.com/pahan-fe/lite-streaming/backend/internal/storage"
)

type VideoService struct {
	repo    *repository.VideoRepository
	queue   *queue.RabbitMQ
	storage *storage.S3Storage
}

func (s *VideoService) Upload(videoData []byte, contentType string, filename string) (string, error) {
	id := uuid.New().String()

	storageErr := s.storage.Upload("videos/"+id+"/original/"+filename, videoData, contentType)
	if storageErr != nil {
		return id, storageErr
	}

	video := &model.Video{
		ID:               id,
		OriginalFilename: filename,
		ContentType:      contentType,
		Size:             int64(len(videoData)),
		Status:           "uploaded",
		S3RawKey:         "videos/" + id + "/original/" + filename,
		S3HLSKey:         "",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	repoErr := s.repo.Create(video)
	if repoErr != nil {
		return id, repoErr
	}

	jsonData, err := json.Marshal(map[string]string{"video_id": id})
	if err != nil {
		return id, err
	}

	queueErr := s.queue.Publish("transcode", jsonData)
	if queueErr != nil {
		return id, queueErr
	}

	return id, nil
}

func (s *VideoService) GetByID(id string) (*model.Video, error) {
	video, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (s *VideoService) Delete(id string) error {
	video, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	deleteErr := s.storage.Delete(video.S3RawKey)
	if deleteErr != nil {
		return deleteErr
	}

	deleteErr = s.repo.Delete(video.ID)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func (s *VideoService) List(page int, limit int) ([]model.Video, error) {
	videos, err := s.repo.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (s *VideoService) GetRawStream(id string) ([]byte, string, error) {
	video, err := s.repo.GetByID(id)
	if err != nil {
		return nil, "", err
	}

	data, storageErr := s.storage.Get(video.S3RawKey)
	if storageErr != nil {
		return nil, "", storageErr
	}

	return data, video.ContentType, nil
}

func (s *VideoService) GetHLSFile(id string, filename string) ([]byte, string, error) {
	_, videoErr := s.repo.GetByID(id)
	if videoErr != nil {
		return nil, "", videoErr
	}

	key := "videos/" + id + "/hls/" + filename
	data, storageErr := s.storage.Get(key)

	if storageErr != nil {
		return nil, "", storageErr
	}

	var contentType string
	if strings.HasSuffix(filename, ".m3u8") {
		contentType = "application/vnd.apple.mpegurl"
	} else {
		contentType = "video/mp2t"
	}

	return data, contentType, nil
}

func NewVideoService(repo *repository.VideoRepository, queue *queue.RabbitMQ, storage *storage.S3Storage) *VideoService {
	return &VideoService{repo: repo, queue: queue, storage: storage}
}
