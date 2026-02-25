package master

import (
	"context"
	"log"

	domainService "github.com/R2Remote/ChronoGo/internal/domain/service"
)

// Start 初始化 Master 节点的核心组件
func Start() {

	ctx := context.Background()

	dispatcher := domainService.NewDispatcher()

	dispatcher.Init(ctx)

	dispatcher.Start(ctx)

	log.Println("Master node initialized and consumer started.")
}
