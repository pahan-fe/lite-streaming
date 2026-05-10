package service

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pahan-fe/lite-streaming/backend/internal/model"
)

type Repo interface {
	Create(ctx context.Context, video *model.Video) error
	GetByID(ctx context.Context, id string) (*model.Video, error)
	GetAll(ctx context.Context, page int, limit int) ([]model.Video, error)
	Delete(ctx context.Context, id string) error
}

type Publisher interface {
	Publish(ctx context.Context, queueName string, body []byte) error
}

type Storage interface {
	Upload(ctx context.Context, key string, data []byte, contentType string) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}

type VideoService struct {
	repo    Repo
	queue   Publisher
	storage Storage
}

func (s *VideoService) Upload(ctx context.Context, videoData []byte, contentType string, filename string) (string, error) {
	id := uuid.New().String()

	storageErr := s.storage.Upload(ctx, "videos/"+id+"/original/"+filename, videoData, contentType)
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
	repoErr := s.repo.Create(ctx, video)
	if repoErr != nil {
		return id, repoErr
	}

	jsonData, err := json.Marshal(map[string]string{"video_id": id})
	if err != nil {
		return id, err
	}

	queueErr := s.queue.Publish(ctx, "transcode", jsonData)
	if queueErr != nil {
		return id, queueErr
	}

	return id, nil
}

func (s *VideoService) GetByID(ctx context.Context, id string) (*model.Video, error) {
	video, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (s *VideoService) Delete(ctx context.Context, id string) error {
	video, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	deleteErr := s.storage.Delete(ctx, video.S3RawKey)
	if deleteErr != nil {
		return deleteErr
	}

	deleteErr = s.repo.Delete(ctx, video.ID)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func (s *VideoService) List(ctx context.Context, page int, limit int) ([]model.Video, error) {
	videos, err := s.repo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (s *VideoService) GetRawStream(ctx context.Context, id string) ([]byte, string, error) {
	video, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, "", err
	}

	data, storageErr := s.storage.Get(ctx, video.S3RawKey)
	if storageErr != nil {
		return nil, "", storageErr
	}

	return data, video.ContentType, nil
}

func (s *VideoService) GetHLSFile(ctx context.Context, id string, filename string) ([]byte, string, error) {
	_, videoErr := s.repo.GetByID(ctx, id)
	if videoErr != nil {
		return nil, "", videoErr
	}

	key := "videos/" + id + "/hls/" + filename
	data, storageErr := s.storage.Get(ctx, key)

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

func NewVideoService(repo Repo, queue Publisher, storage Storage) *VideoService {
	return &VideoService{repo: repo, queue: queue, storage: storage}
}
