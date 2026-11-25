package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pahan-fe/lite-streaming/backend/internal/model"
)

type Repo interface {
	Create(ctx context.Context, video *model.Video) error
	GetByID(ctx context.Context, id string) (*model.Video, error)
	GetAll(ctx context.Context, cursor *model.Cursor, limit int) ([]model.Video, error)
	Delete(ctx context.Context, id string) error
}

type Publisher interface {
	Publish(ctx context.Context, queueName string, body []byte) error
}

type Storage interface {
	Upload(ctx context.Context, key string, reader io.Reader, contentType string, size int64) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}

type VideoService struct {
	repo    Repo
	queue   Publisher
	storage Storage
}

func encodeCursor(c model.Cursor) string {
	raw := c.CreatedAt.Format(time.RFC3339Nano) + "|" + c.ID
	return base64.URLEncoding.EncodeToString([]byte(raw))
}

func decodeCursor(s string) (*model.Cursor, error) {
	if s == "" {
		return nil, nil
	}

	decoded, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	parts := strings.SplitN(string(decoded), "|", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid cursor")
	}

	createdAt, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil {
		return nil, err
	}

	return &model.Cursor{ID: parts[1], CreatedAt: createdAt}, nil
}

func (s *VideoService) Upload(ctx context.Context, reader io.Reader, contentType string, filename string, fileSize int64) (string, error) {
	id := uuid.New().String()

	storageErr := s.storage.Upload(ctx, "videos/"+id+"/original/"+filename, reader, contentType, fileSize)
	if storageErr != nil {
		return id, storageErr
	}

	video := &model.Video{
		ID:               id,
		OriginalFilename: filename,
		ContentType:      contentType,
		Size:             fileSize,
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

func (s *VideoService) List(ctx context.Context, cursor string, limit int) (model.VideoListResponse, error) {
	decodedCursor, decodeErr := decodeCursor(cursor)
	if decodeErr != nil {
		return model.VideoListResponse{}, decodeErr
	}

	videos, err := s.repo.GetAll(ctx, decodedCursor, limit+1)
	if err != nil {
		return model.VideoListResponse{}, err
	}

	hasNext := len(videos) > limit
	if hasNext {
		videos = videos[:limit]
	}

	var nextCursor *string
	if hasNext {
		lastVideo := videos[len(videos)-1]
		encoded := encodeCursor(model.Cursor{ID: lastVideo.ID, CreatedAt: lastVideo.CreatedAt})
		nextCursor = &encoded
	}

	return model.VideoListResponse{Items: videos, NextCursor: nextCursor}, nil
}

func (s *VideoService) GetRawStream(ctx context.Context, id string) (io.ReadCloser, string, error) {
	video, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, "", err
	}

	reader, storageErr := s.storage.Get(ctx, video.S3RawKey)
	if storageErr != nil {
		return nil, "", storageErr
	}

	return reader, video.ContentType, nil
}

func (s *VideoService) GetHLSFile(ctx context.Context, id string, filename string) (io.ReadCloser, string, error) {
	_, videoErr := s.repo.GetByID(ctx, id)
	if videoErr != nil {
		return nil, "", videoErr
	}

	key := "videos/" + id + "/hls/" + filename
	reader, storageErr := s.storage.Get(ctx, key)

	if storageErr != nil {
		return nil, "", storageErr
	}

	var contentType string
	if strings.HasSuffix(filename, ".m3u8") {
		contentType = "application/vnd.apple.mpegurl"
	} else {
		contentType = "video/mp2t"
	}

	return reader, contentType, nil
}

func NewVideoService(repo Repo, queue Publisher, storage Storage) *VideoService {
	return &VideoService{repo: repo, queue: queue, storage: storage}
}
