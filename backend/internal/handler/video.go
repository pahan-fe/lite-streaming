package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/pahan-fe/lite-streaming/backend/internal/service"
)

type VideoHandler struct {
	service *service.VideoService
}

func (h *VideoHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 30)
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

	data, readErr := io.ReadAll(file)
	if readErr != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	id, uploadErr := h.service.Upload(data, header.Header.Get("Content-Type"), header.Filename)
	if uploadErr != nil {
		http.Error(w, "Error uploading video", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *VideoHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, pageErr := strconv.Atoi(pageStr)
	if pageErr != nil || page < 1 {
		page = 1
	}

	limit, limitErr := strconv.Atoi(limitStr)
	if limitErr != nil || limit < 1 {
		limit = 20
	}

	videos, videosErr := h.service.List(page, limit)
	if videosErr != nil {
		http.Error(w, "Error fetching videos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}

func (h *VideoHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	video, videoErr := h.service.GetByID(id)
	if videoErr != nil {
		http.Error(w, "Video not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
}

func (h *VideoHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	deleteErr := h.service.Delete(id)
	if deleteErr != nil {
		http.Error(w, "Error deleting video", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *VideoHandler) HandleStream(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	data, contentType, err := h.service.GetRawStream(id)
	if err != nil {
		http.Error(w, "Error fetching video stream", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(data)
}

func (h *VideoHandler) HandleHLSFile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	filename := r.PathValue("filename")

	data, contentType, err := h.service.GetHLSFile(id, filename)
	if err != nil {
		http.Error(w, "Error fetching HLS file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(data)
}

func NewVideoHandler(service *service.VideoService) *VideoHandler {
	return &VideoHandler{service: service}
}
