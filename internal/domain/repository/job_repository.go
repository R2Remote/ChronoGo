package repository

import (
	"context"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
)

type JobRepository interface {
	Create(ctx context.Context, job *entity.Job) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (*entity.Job, error)
	List(ctx context.Context, page, pageSize int) ([]*entity.Job, int64, error)
	Update(ctx context.Context, job *entity.Job) error
}
