package service

import (
	"context"
	"log"
	"math"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/R2Remote/ChronoGo/internal/domain/repository"
)

type Dispatcher struct {
	configMap map[uint64]entity.Job
	serverMap map[uint64]entity.Server
	repo      repository.JobRepository
}

func (d *Dispatcher) FetchServer() {

}

func NewDispatcher(repo repository.JobRepository) *Dispatcher {
	return &Dispatcher{
		configMap: make(map[uint64]entity.Job),
		repo:      repo,
	}
}

func (d *Dispatcher) LoadConfig(ctx context.Context) error {
	jobs, _, err := d.repo.List(ctx, 1, math.MaxInt32)
	if err != nil {
		return err
	}

	for _, job := range jobs {
		log.Println("Load job:", job.ID, job.Name)
		d.configMap[job.ID] = *job
	}
	return nil
}

func (d *Dispatcher) Dispatch(task *entity.Task) {
	job, ok := d.configMap[task.JobID]

	if !ok {
		log.Printf("Job %d not found in config", task.JobID)
		return
	}

	// TODO: Implement actual dispatch logic
	log.Printf("Dispatching job: %s (ID: %d)", job.Name, job.ID)
}
