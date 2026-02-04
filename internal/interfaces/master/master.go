package master

import (
	"context"
	"log"
	"math"
	"strconv"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/database"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/redis"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/repository"
)

type JobConfig struct {
	configMap map[uint64]entity.Job
}

var jobConfig *JobConfig

var jobQueue chan uint64

func ListenAndDispatch() {
	select {
	case jobId := <-jobQueue:
		job := jobConfig.configMap[jobId]
		dispatch(&job)
	}
}

func dispatch(job *entity.Job) {

}

// 从redis中取任务
func scanJobFromRedis() {
	for {
		// BLPop returns []string{key, value}
		result, err := redis.Client.BLPop(context.Background(), 0, "job_queue").Result()
		if err != nil {
			log.Println("get job from redis error:", err)
			continue
		}
		if len(result) < 2 {
			continue
		}

		val := result[1]
		jobId, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			log.Printf("invalid job id: %s, error: %v", val, err)
			continue
		}

		jobQueue <- jobId
	}
}

// Init 使用显式初始化代替 init()，因为数据库和Redis连接在 main 中才初始化
func Init() {
	jobQueue = make(chan uint64, 100) // 带缓冲

	// load job config
	loadJobConfig()

	// start get job from redis
	go scanJobFromRedis()

	// start consumer
	go consumeJobs()
}

func consumeJobs() {
	for jobId := range jobQueue {
		if job, ok := jobConfig.configMap[jobId]; ok {
			dispatch(&job)
		} else {
			log.Printf("job %d not found in config", jobId)
		}
	}
}

func loadJobConfig() {
	jobConfig = &JobConfig{
		configMap: make(map[uint64]entity.Job),
	}
	repo := repository.NewJobRepository(database.DB)
	jobs, _, err := repo.List(context.Background(), 1, math.MaxInt32)
	if err != nil {
		panic(err)
	}
	for _, job := range jobs {
		log.Println("load job:", job.ServerID, job.Name)
		jobConfig.configMap[job.ID] = *job
	}
}
