package service

import (
	"github.com/pahan-fe/lite-streaming/backend/internal/repository"
	"github.com/pahan-fe/lite-streaming/backend/internal/storage"
	"github.com/pahan-fe/lite-streaming/backend/internal/queue"
	"github.com/pahan-fe/lite-streaming/backend/internal/model"
	"github.com/google/uuid"  
	"time"
	"encoding/json"
)

type VideoService struct {
	repo *repository.VideoRepository
	queue *queue.RabbitMQ
	storage *storage.S3Storage
}

func (s *VideoService) Upload(videoData []byte, contentType string, filename string) (string, error) {
	id := uuid.New().String()

	storageErr := s.storage.Upload("videos/" + id + "/original/" + filename, videoData, contentType)
	if storageErr != nil {
		return id, storageErr
	}
	
	video := &model.Video{
		ID: id,
		OriginalFilename: filename,
		ContentType: contentType,
		Size: int64(len(videoData)),
		Status: "uploaded",
		S3RawKey: "videos/" + id + "/original/" + filename,
		S3HLSKey: "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

func (s *VideoService) List(page int, limit int) ([]*model.Video, error) {
	videos, err := s.repo.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func NewService(repo *repository.VideoRepository, queue *queue.RabbitMQ, storage *storage.S3Storage) *VideoService {
	return &VideoService{repo: repo, queue: queue, storage: storage}
}