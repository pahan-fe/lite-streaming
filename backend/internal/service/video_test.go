package service

import (
	"context"
	"errors"
	"testing"

	"github.com/pahan-fe/lite-streaming/backend/internal/model"
)

type mockRepo struct {
	createCalled bool
	createErr    error
	getByIDVideo *model.Video
	getByIDErr   error
	deleteCalled bool
	deleteErr    error
}

type mockPublisher struct {
	publishCalled bool
	publishErr    error
}

type mockStorage struct {
	uploadErr error
	deleteErr error
}

func (m *mockRepo) Create(ctx context.Context, video *model.Video) error {
	m.createCalled = true
	return m.createErr
}
func (m *mockRepo) GetByID(ctx context.Context, id string) (*model.Video, error) {

	return m.getByIDVideo, m.getByIDErr
}
func (m *mockRepo) GetAll(ctx context.Context, page int, limit int) ([]model.Video, error) {
	return nil, nil
}
func (m *mockRepo) Delete(ctx context.Context, id string) error {
	m.deleteCalled = true
	return m.deleteErr
}

func (m *mockPublisher) Publish(ctx context.Context, queueName string, body []byte) error {
	m.publishCalled = true
	return m.publishErr
}

func (m *mockStorage) Get(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}
func (m *mockStorage) Upload(ctx context.Context, key string, data []byte, contentType string) error {
	return m.uploadErr
}
func (m *mockStorage) Delete(ctx context.Context, key string) error {
	return m.deleteErr
}

func TestVideoService_Upload_HappyPath(t *testing.T) {
	repo := &mockRepo{}
	queue := &mockPublisher{}
	storage := &mockStorage{}
	svc := NewVideoService(repo, queue, storage)

	id, err := svc.Upload(context.Background(), []byte("video data"), "video/mp4", "test.mp4")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id == "" {
		t.Fatalf("expected a valid video ID, got empty string")
	}
}

func TestVideoService_Upload_StorageError(t *testing.T) {
	repo := &mockRepo{}
	queue := &mockPublisher{}
	storage := &mockStorage{uploadErr: errors.New("s3 down")}
	svc := NewVideoService(repo, queue, storage)

	_, err := svc.Upload(context.Background(), []byte("data"), "video/mp4", "test.mp4")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if repo.createCalled {
		t.Error("expected repo.Create not to be called when storage upload fails")
	}
	if queue.publishCalled {
		t.Error("expected queue.Publish not to be called when storage upload fails")
	}
}

func TestVideoService_Upload_RepoError(t *testing.T) {
	repo := &mockRepo{createErr: errors.New("db down")}
	queue := &mockPublisher{}
	storage := &mockStorage{}
	svc := NewVideoService(repo, queue, storage)

	_, err := svc.Upload(context.Background(), []byte("data"), "video/mp4", "test.mp4")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if queue.publishCalled {
		t.Error("expected queue.Publish not to be called when repo.Create fails")
	}
}

func TestVideoService_Upload_QueueError(t *testing.T) {
	repo := &mockRepo{}
	queue := &mockPublisher{publishErr: errors.New("rabbitmq down")}
	storage := &mockStorage{}
	svc := NewVideoService(repo, queue, storage)

	_, err := svc.Upload(context.Background(), []byte("data"), "video/mp4", "test.mp4")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestVideoService_Delete_HappyPath(t *testing.T) {
	repo := &mockRepo{getByIDVideo: &model.Video{ID: "123", S3RawKey: "raw/123.mp4"}}
	storage := &mockStorage{}
	queue := &mockPublisher{}
	svc := NewVideoService(repo, queue, storage)

	err := svc.Delete(context.Background(), "123")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestVideoService_Delete_GetByIDError(t *testing.T) {
	repo := &mockRepo{getByIDVideo: nil, getByIDErr: errors.New("video not found")}
	storage := &mockStorage{}
	queue := &mockPublisher{}
	svc := NewVideoService(repo, queue, storage)

	err := svc.Delete(context.Background(), "123")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if repo.deleteCalled {
		t.Error("expected repo.Delete not to be called when repo.GetByID fails")
	}
}

func TestVideoService_Delete_StorageError(t *testing.T) {
	repo := &mockRepo{getByIDVideo: &model.Video{ID: "123", S3RawKey: "raw/123.mp4"}}
	storage := &mockStorage{deleteErr: errors.New("s3 down")}
	queue := &mockPublisher{}
	svc := NewVideoService(repo, queue, storage)

	err := svc.Delete(context.Background(), "123")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if repo.deleteCalled {
		t.Error("expected repo.Delete NOT to be called when storage.Delete fails")
	}
}

func TestVideoService_Delete_RepoError(t *testing.T) {
	repo := &mockRepo{getByIDVideo: &model.Video{ID: "123", S3RawKey: "raw/123.mp4"}, deleteErr: errors.New("db down")}
	storage := &mockStorage{}
	queue := &mockPublisher{}
	svc := NewVideoService(repo, queue, storage)

	err := svc.Delete(context.Background(), "123")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !repo.deleteCalled {
		t.Error("expected repo.Delete to be called when storage.Delete succeeds")
	}
}
