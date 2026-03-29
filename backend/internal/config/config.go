package config

import (
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL     string
	S3Endpoint      string
	S3AccessKey     string
	S3SecretKey     string
	S3Bucket        string
	RabbitMQURL     string
	S3UseSSL        bool
	MaxUploadSize   int64

}

func Load() Config {
	s3UseSsl, _ := strconv.ParseBool(os.Getenv("S3_USE_SSL"))
	maxUploadSize, _ := strconv.ParseInt(os.Getenv("MAX_UPLOAD_SIZE"), 10, 64)
	
	return Config{
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		S3Endpoint:      os.Getenv("S3_ENDPOINT"),
		S3AccessKey:     os.Getenv("S3_ACCESS_KEY"),
		S3SecretKey:     os.Getenv("S3_SECRET_KEY"),
		S3Bucket:        os.Getenv("S3_BUCKET"),
		RabbitMQURL:     os.Getenv("RABBITMQ_URL"),
		S3UseSSL:        s3UseSsl,
		MaxUploadSize:   maxUploadSize,
	}
}