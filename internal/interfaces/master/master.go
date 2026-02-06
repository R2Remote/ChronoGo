package master

import (
	"context"
	"log"

	domainService "github.com/R2Remote/ChronoGo/internal/domain/service"
)

// Init 初始化 Master 节点的核心组件
func Init() {

	ctx := context.Background()

	dispatcher := domainService.NewDispatcher()

	dispatcher.FetchJob(ctx)

	dispatcher.FetchServer(ctx)

	dispatcher.Start(ctx)

	log.Println("Master node initialized and consumer started.")
}
