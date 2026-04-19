package storage

import (
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pahan-fe/lite-streaming/backend/internal/config"
)

type S3Storage struct {
	client *minio.Client
	bucket string
}

func (s *S3Storage) Upload(key string, data []byte, contentType string) error {
	reader := bytes.NewReader(data)

	_, err := s.client.PutObject(context.Background(), s.bucket, key, reader, int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	return nil
}

func (s *S3Storage) Get(key string) ([]byte, error) {
	obj, objErr := s.client.GetObject(context.Background(), s.bucket, key, minio.GetObjectOptions{})
	if objErr != nil {
		return nil, objErr
	}

	data, convertErr := io.ReadAll(obj)
	if convertErr != nil {
		return nil, convertErr
	}

	return data, nil
}

func (s *S3Storage) Delete(key string) error {
	err := s.client.RemoveObject(context.Background(), s.bucket, key, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func NewS3Storage(c *config.Config) (*S3Storage, error) {
	s, sErr := minio.New(c.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.S3AccessKey, c.S3SecretKey, ""),
		Secure: c.S3UseSSL,
	})
	if sErr != nil {
		return nil, sErr
	}

	ctx := context.Background()

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
