package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/pahan-fe/lite-streaming/backend/internal/service"
)

type VideoHandler struct {
	service *service.VideoService
}

func (h *VideoHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	id, uploadErr := h.service.Upload(ctx, file, header.Header.Get("Content-Type"), header.Filename, header.Size)
	if uploadErr != nil {
		http.Error(w, "Error uploading video", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *VideoHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cursor := r.URL.Query().Get("cursor")
	limitStr := r.URL.Query().Get("limit")

	limit, limitErr := strconv.Atoi(limitStr)
	if limitErr != nil || limit < 1 {
		limit = 20
	}

	videos, videosErr := h.service.List(ctx, cursor, limit)
	if videosErr != nil {
		http.Error(w, "Error fetching videos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}

func (h *VideoHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")

	video, videoErr := h.service.GetByID(ctx, id)
	if videoErr != nil {
		http.Error(w, "Video not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
}

func (h *VideoHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")

	deleteErr := h.service.Delete(ctx, id)
	if deleteErr != nil {
		http.Error(w, "Error deleting video", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *VideoHandler) HandleStream(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")

	reader, contentType, err := h.service.GetRawStream(ctx, id)
	if err != nil {
		http.Error(w, "Error fetching video stream", http.StatusInternalServerError)
		return
	}
	defer reader.Close()

	w.Header().Set("Content-Type", contentType)
	_, writingErr := io.Copy(w, reader)
	if writingErr != nil {
		slog.ErrorContext(ctx, "Failed writing to response", "err", writingErr)
	}
}

func (h *VideoHandler) HandleHLSFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	filename := r.PathValue("filename")

	reader, contentType, err := h.service.GetHLSFile(ctx, id, filename)
	if err != nil {
		http.Error(w, "Error fetching HLS file", http.StatusInternalServerError)
		return
	}
	defer reader.Close()

	w.Header().Set("Content-Type", contentType)
	_, copyErr := io.Copy(w, reader)
	if copyErr != nil {
		slog.ErrorContext(ctx, "Failed splitting file", "err", copyErr)
	}
}

func NewVideoHandler(service *service.VideoService) *VideoHandler {
	return &VideoHandler{service: service}
}
