package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pahan-fe/lite-streaming/backend/internal/config"
	"github.com/pahan-fe/lite-streaming/backend/internal/handler"
	"github.com/pahan-fe/lite-streaming/backend/internal/middleware"
	"github.com/pahan-fe/lite-streaming/backend/internal/queue"
	"github.com/pahan-fe/lite-streaming/backend/internal/repository"
	"github.com/pahan-fe/lite-streaming/backend/internal/service"
	"github.com/pahan-fe/lite-streaming/backend/internal/storage"
)

func main() {
	base := slog.NewJSONHandler(os.Stdout, nil)
	slogHandler := &middleware.ContextHandler{Handler: base}
	slog.SetDefault(slog.New(slogHandler))

	mux := http.NewServeMux()
	loggerMux := middleware.RequestID(middleware.Logging(middleware.Recovery(mux)))

	slog.Info("starting api server", "port", 8080)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("POST /api/videos", videoHandler.HandleUpload)
	mux.HandleFunc("GET /api/videos", videoHandler.HandleList)
	mux.HandleFunc("GET /api/videos/{id}", videoHandler.HandleGetByID)
	mux.HandleFunc("DELETE /api/videos/{id}", videoHandler.HandleDelete)
	mux.HandleFunc("GET /api/videos/{id}/stream", videoHandler.HandleStream)
	mux.HandleFunc("GET /api/videos/{id}/hls/{filename}", videoHandler.HandleHLSFile)

	err := http.ListenAndServe(":8080", loggerMux)
	if err != nil {
		slog.Error("Failed to start server", "err", err)
		os.Exit(1)
	}
}
