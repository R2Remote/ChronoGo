package service

import (
	"context"
	"log"
	"math"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/database"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/redisgo"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/repository"
)

type Dispatcher struct {
	configMap map[uint64]entity.Job
	serverMap map[string]entity.Server
}

func (d *Dispatcher) Start(ctx context.Context) {
	consumer := NewTaskConsumer(d, redisgo.Client, ctx)
	go d.loop(consumer)
}

func (d *Dispatcher) loop(consumer *TaskConsumer) {
	for {
		task := consumer.FetchTask()
		d.Dispatch(task)
	}
}

func (d *Dispatcher) FetchJob(ctx context.Context) {
	infraRepo := repository.NewJobRepository(database.DB)
	jobService := NewJobService(infraRepo)
	jobs, _ := jobService.List(ctx, 1, math.MaxInt32)
	for _, job := range jobs {
		log.Println("Load job:", job.ID, job.Name)
		d.configMap[job.ID] = *job
	}
	go jobService.Loop(&d.configMap)
}

func (d *Dispatcher) FetchServer(ctx context.Context) {
	serverService := NewServerService(redisgo.Client)
	servers, _ := serverService.List(ctx, 1, math.MaxInt32)
	for _, server := range servers {
		log.Println("Load server:", server.ID, server.Name)
		d.serverMap[server.Name] = *server
	}
	go serverService.Loop(&d.serverMap)
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		configMap: make(map[uint64]entity.Job),
		serverMap: make(map[string]entity.Server),
	}
}

func (d *Dispatcher) Dispatch(task *entity.Task) {
	job, ok := d.configMap[task.JobID]

	if !ok {
		log.Printf("Job %d not found in config", task.JobID)
		return
	}
	serverName := job.ServerName
	server, ok := d.serverMap[serverName]
	if !ok {
		log.Printf("Server %d not found in config", serverName)
		return
	}
	//发送RPC调用

	// TODO: Implement actual dispatch logic
	log.Printf("Dispatching job: %s (ID: %d) to server: %s (ID: %d)", job.Name, job.ID, server.Name, server.ID)
}
