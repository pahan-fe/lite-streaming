package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pahan-fe/lite-streaming/backend/internal/config"
	"github.com/pahan-fe/lite-streaming/backend/internal/handler"
	"github.com/pahan-fe/lite-streaming/backend/internal/queue"
	"github.com/pahan-fe/lite-streaming/backend/internal/repository"
	"github.com/pahan-fe/lite-streaming/backend/internal/service"
	"github.com/pahan-fe/lite-streaming/backend/internal/storage"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("starting api server", "port", 8080)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	cfg := config.Load()
	db, dbErr := sqlx.Connect("postgres", cfg.DatabaseURL)
	if dbErr != nil {
		slog.Error("Failed to connect to database", "err", dbErr)
		os.Exit(1)
	}
	defer db.Close()

	str, storageErr := storage.NewS3Storage(&cfg)
	if storageErr != nil {
		slog.Error("Failed to initialize storage", "err", storageErr)
		os.Exit(1)
	}

	mq, queueErr := queue.NewRabbitMQ(&cfg)
	if queueErr != nil {
		slog.Error("Failed to initialize queue", "err", queueErr)
		os.Exit(1)
	}

	repo := repository.NewVideoRepository(db)
	videoService := service.NewVideoService(repo, mq, str)
	videoHandler := handler.NewVideoHandler(videoService)

	http.HandleFunc("POST /api/videos", videoHandler.HandleUpload)
	http.HandleFunc("GET /api/videos", videoHandler.HandleList)
	http.HandleFunc("GET /api/videos/{id}", videoHandler.HandleGetByID)
	http.HandleFunc("DELETE /api/videos/{id}", videoHandler.HandleDelete)
	http.HandleFunc("GET /api/videos/{id}/stream", videoHandler.HandleStream)
	http.HandleFunc("GET /api/videos/{id}/hls/{filename}", videoHandler.HandleHLSFile)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("Failed to start server", "err", err)
		os.Exit(1)
	}
}
