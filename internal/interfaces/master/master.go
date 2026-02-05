package master

import (
	"context"
	"log"

	applicationService "github.com/R2Remote/ChronoGo/internal/application/service"
	domainService "github.com/R2Remote/ChronoGo/internal/domain/service"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/database"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/redis"
	infraRepo "github.com/R2Remote/ChronoGo/internal/infrastructure/repository"
)

// Init 初始化 Master 节点的核心组件
func Init() {
	// 1. 初始化 Repository (Infrastructure Layer)
	jobRepo := infraRepo.NewJobRepository(database.DB)

	// 2. 初始化 Dispatcher (Domain Service Layer)
	dispatcher := domainService.NewDispatcher(jobRepo)
	if err := dispatcher.LoadConfig(context.Background()); err != nil {
		log.Fatalf("Failed to load job config: %v", err)
	}
	dispatcher.FetchServer()

	// 3. 初始化 Consumer (Application Service Layer)
	consumer := applicationService.NewTaskConsumer(dispatcher, redis.Client)

	// 4. 启动消费者
	consumer.Start()

	log.Println("Master node initialized and consumer started.")
}
