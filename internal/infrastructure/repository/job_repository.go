package repository

import (
	"context"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/dao"
	"gorm.io/gorm"
)

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *jobRepository {
	return &jobRepository{db: db}
}

func (r *jobRepository) Create(ctx context.Context, job *entity.Job) error {
	po := r.toPO(job)
	if err := r.db.WithContext(ctx).Create(po).Error; err != nil {
		return err
	}
	job.ID = po.ID
	job.CreatedAt = po.CreatedAt
	job.UpdatedAt = po.UpdatedAt
	return nil
}

func (r *jobRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&dao.Job{}, id).Error
}

func (r *jobRepository) FindByID(ctx context.Context, id uint64) (*entity.Job, error) {
	var po dao.Job
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&po).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.toEntity(&po), nil
}

func (r *jobRepository) List(ctx context.Context, page, pageSize int) ([]*entity.Job, int64, error) {
	var pos []*dao.Job
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.WithContext(ctx).Model(&dao.Job{}).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&pos).Error; err != nil {
		return nil, 0, err
	}

	jobs := make([]*entity.Job, len(pos))
	for i, po := range pos {
		jobs[i] = r.toEntity(po)
	}

	return jobs, total, nil
}

func (r *jobRepository) Update(ctx context.Context, job *entity.Job) error {
	po := r.toPO(job)
	return r.db.WithContext(ctx).Model(&dao.Job{ID: job.ID}).Updates(po).Error
}

func (r *jobRepository) toPO(job *entity.Job) *dao.Job {
	return &dao.Job{
		ID:          job.ID,
		CreatedAt:   job.CreatedAt,
		UpdatedAt:   job.UpdatedAt,
		Name:        job.Name,
		Description: job.Description,
		Command:     job.Command,
		CronSpec:    job.CronSpec,
		RetryCount:  job.RetryCount,
		Timeout:     job.Timeout,
	}
}

func (r *jobRepository) toEntity(po *dao.Job) *entity.Job {
	return &entity.Job{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		Name:        po.Name,
		Description: po.Description,
		Command:     po.Command,
		CronSpec:    po.CronSpec,
		RetryCount:  po.RetryCount,
		Timeout:     po.Timeout,
	}
}
