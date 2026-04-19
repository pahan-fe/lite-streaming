package main

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Println("Starting API server on port 8080...")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	cfg := config.Load()
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	str, storageErr := storage.NewS3Storage(&cfg)
	if storageErr != nil {
		log.Fatalf("Failed to initialize storage: %v", storageErr)
	}

	mq, queueErr := queue.NewRabbitMQ(&cfg)
	if queueErr != nil {
		log.Fatalf("Failed to initialize queue: %v", queueErr)
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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
