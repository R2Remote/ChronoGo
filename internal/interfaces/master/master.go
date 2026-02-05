package master

import (
	"context"
	"log"

	applicationService "github.com/R2Remote/ChronoGo/internal/application/service"
	domainService "github.com/R2Remote/ChronoGo/internal/domain/service"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/redis"
)

// Init 初始化 Master 节点的核心组件
func Init() {

	ctx := context.Background()

	dispatcher := domainService.NewDispatcher()

	dispatcher.FetchJob(ctx)

	dispatcher.FetchServer(ctx)

	// 3. 初始化 Consumer (Application Service Layer)
	consumer := applicationService.NewTaskConsumer(dispatcher, redis.Client)

	// 4. 启动消费者
	consumer.Start()

	log.Println("Master node initialized and consumer started.")
}
