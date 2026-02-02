package master

import (
	"context"
	"log"
	"math"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/database"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/repository"
)

type JobCache struct {
	jobMap map[uint64]map[uint64]entity.Job
}

var jobCache *JobCache

func ListenAndDispatch() {
	log.Println("todo")
}

func init() {
	jobCache = &JobCache{
		jobMap: make(map[uint64]map[uint64]entity.Job),
	}
	repo := repository.NewJobRepository(database.DB)
	jobs, _, err := repo.List(context.Background(), 1, math.MaxInt32)
	if err != nil {
		panic(err)
	}
	for _, job := range jobs {
		log.Println("load job:", job.ServerID, job.Name)
		jobServerMap, exists := jobCache.jobMap[job.ServerID]
		if !exists {
			jobServerMap = make(map[uint64]entity.Job)
			jobCache.jobMap[job.ServerID] = jobServerMap
		}
		jobServerMap[job.ID] = *job
	}
}
