package service

import (
	"context"
	"log"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/R2Remote/ChronoGo/internal/domain/repository"
)

type JobService struct {
	repo repository.JobRepository
}

func NewJobService(repo repository.JobRepository) *JobService {
	return &JobService{repo: repo}
}

func (j *JobService) List(ctx context.Context, page, pageSize int) ([]*entity.Job, int64) {
	jobs, c, err := j.repo.List(ctx, page, pageSize)
	if err != nil {
		log.Println("Failed to list jobs:", err)
		return nil, 0
	}
	return jobs, c
}
