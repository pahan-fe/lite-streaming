package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pahan-fe/lite-streaming/backend/internal/model"
)

type VideoRepository struct {
	db *sqlx.DB
}

func (r *VideoRepository) Create(video *model.Video) error {
	_, err := r.db.NamedExec(`
		INSERT INTO videos (id, original_filename, content_type, size, status, s3_raw_key, s3_hls_key, created_at, updated_at) 
		VALUES (:id, :original_filename, :content_type, :size, :status, :s3_raw_key, :s3_hls_key, :created_at, :updated_at)`, video)
	return err
}

func (r *VideoRepository) GetByID(id string) (*model.Video, error) {
	var video model.Video
	err := r.db.Get(&video, "SELECT * FROM videos WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (r *VideoRepository) GetAll(page int, limit int) ([]model.Video, error) {
	videos := []model.Video{}

	offset := (page - 1) * limit

	err := r.db.Select(&videos, "SELECT * FROM videos ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *VideoRepository) UpdateStatus(id string, status string) error {
	_, err := r.db.Exec("UPDATE videos SET status = $1, updated_at = NOW() WHERE id = $2", status, id)
	return err
}

func (r *VideoRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM videos WHERE id = $1", id)
	return err
}

func NewVideoRepository(db *sqlx.DB) *VideoRepository {
	return &VideoRepository{db: db}
}
