package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pahan-fe/lite-streaming/backend/internal/model"
)

type VideoRepository struct {
	db *sqlx.DB
}

func (r *VideoRepository) Create(ctx context.Context, video *model.Video) error {
	_, err := r.db.NamedExecContext(ctx, `
		INSERT INTO videos (id, original_filename, content_type, size, status, s3_raw_key, s3_hls_key, created_at, updated_at) 
		VALUES (:id, :original_filename, :content_type, :size, :status, :s3_raw_key, :s3_hls_key, :created_at, :updated_at)`, video)
	return err
}

func (r *VideoRepository) GetByID(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video
	err := r.db.GetContext(ctx, &video, "SELECT * FROM videos WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (r *VideoRepository) GetAll(ctx context.Context, page int, limit int) ([]model.Video, error) {
	videos := []model.Video{}

	offset := (page - 1) * limit

	err := r.db.SelectContext(ctx, &videos, "SELECT * FROM videos ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *VideoRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE videos SET status = $1, updated_at = NOW() WHERE id = $2", status, id)
	return err
}

func (r *VideoRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM videos WHERE id = $1", id)
	return err
}

func NewVideoRepository(db *sqlx.DB) *VideoRepository {
	return &VideoRepository{db: db}
}
