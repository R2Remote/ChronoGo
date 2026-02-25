package service

import (
	"context"
	"log"
	"math"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/R2Remote/ChronoGo/internal/domain/service/strategy"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/database"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/redisgo"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/repository"
)

type Dispatcher struct {
	jobMap     map[uint64]entity.Job
	serverMap  map[string]entity.Server
	strategies map[entity.StrategyType]strategy.DispatchStrategy
}

func (d *Dispatcher) Init(ctx context.Context) {

	// fetch all Job config
	d.fetchJob(ctx)

	// fetch all alive server
	d.fetchServer(ctx)

	// init strategies
	d.initStrategies()
}

func (d *Dispatcher) initStrategies() {
	d.strategies = make(map[entity.StrategyType]strategy.DispatchStrategy)
	d.strategies[entity.Random] = strategy.NewRandomStrategy()
	d.strategies[entity.RoundRobin] = strategy.NewRoundRobinStrategy()
	d.strategies[entity.Broadcast] = strategy.NewBroadcastStrategy()
	d.strategies[entity.Specific] = strategy.NewSpecificStrategy()
}

func (d *Dispatcher) Start(ctx context.Context) {
	go d.consume(ctx)
}

func (d *Dispatcher) consume(ctx context.Context) {
	consumer := NewTaskConsumer(d, redisgo.Client, ctx)
	go d.loop(consumer)
}

func (d *Dispatcher) loop(consumer *TaskConsumer) {
	for {
		task := consumer.FetchTask()
		d.dispatch(task)
	}
}

func (d *Dispatcher) fetchJob(ctx context.Context) {
	infraRepo := repository.NewJobRepository(database.DB)
	jobService := NewJobService(infraRepo)
	jobs, _ := jobService.List(ctx, 1, math.MaxInt32)
	for _, job := range jobs {
		log.Println("Load job:", job.ID, job.Name)
		d.jobMap[job.ID] = *job
	}
}

func (d *Dispatcher) fetchServer(ctx context.Context) {
	serverService := NewServerService(redisgo.Client)
	servers, _ := serverService.List(ctx, 1, math.MaxInt32)
	for _, server := range servers {
		log.Println("Load server:", server.ID, server.Name)
		d.serverMap[server.Name] = *server
	}
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		jobMap:    make(map[uint64]entity.Job),
		serverMap: make(map[string]entity.Server),
	}
}

func (d *Dispatcher) dispatch(task *entity.Task) {
	job, ok := d.jobMap[task.JobID]
	if !ok {
		log.Printf("Job %d not found in config", task.JobID)
		return
	}
	strategy, exists := d.strategies[task.Strategy]
	if !exists {
		log.Printf("Strategy %d not found in config", task.Strategy)
		return
	}
	servers, err := strategy.SelectServers(&job)
	if err != nil {
		log.Printf("Failed to select nodes: %v", err)
		return
	}
	for _, server := range servers {
		//发送RPC调用
		log.Printf("Dispatching job: %s (ID: %d) to server: %s (ID: %d)", job.Name, job.ID, server.Name, server.ID)
	}
}
