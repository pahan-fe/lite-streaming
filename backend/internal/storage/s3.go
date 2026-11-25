package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pahan-fe/lite-streaming/backend/internal/config"
)

type S3Storage struct {
	client *minio.Client
	bucket string
}

func (s *S3Storage) Upload(ctx context.Context, key string, reader io.Reader, contentType string, fileSize int64) error {
	_, err := s.client.PutObject(ctx, s.bucket, key, reader, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	return nil
}

func (s *S3Storage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	obj, objErr := s.client.GetObject(ctx, s.bucket, key, minio.GetObjectOptions{})
	if objErr != nil {
		return nil, objErr
	}

	_, infoErr := obj.Stat()
	if infoErr != nil {
		obj.Close()
		return nil, fmt.Errorf("Object not found: %v", infoErr)
	}

	return obj, nil
}

func (s *S3Storage) Delete(ctx context.Context, key string) error {
	err := s.client.RemoveObject(ctx, s.bucket, key, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func NewS3Storage(c *config.Config) (*S3Storage, error) {
	ctx := context.Background()

	s, sErr := minio.New(c.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.S3AccessKey, c.S3SecretKey, ""),
		Secure: c.S3UseSSL,
	})
	if sErr != nil {
		return nil, sErr
	}

	exist, existErr := s.BucketExists(ctx, c.S3Bucket)
	if existErr != nil {
		return nil, existErr
	}
	if !exist {
		createErr := s.MakeBucket(ctx, c.S3Bucket, minio.MakeBucketOptions{})
		if createErr != nil {
			return nil, createErr
		}
	}

	return &S3Storage{client: s, bucket: c.S3Bucket}, nil
}
