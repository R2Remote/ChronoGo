package main

import (
	"fmt"
	"log"

	"github.com/R2Remote/ChronoGo/internal/interfaces/api"
)

func main() {

	fmt.Println("🚀 Master 节点启动成功！正在监听任务队列...")

	// 启动 Web API 服务 (对接 FlowBoard)
	apiServer := api.NewServer()
	if err := apiServer.Start(":8080"); err != nil {
		log.Fatalf("Master Web API 启动失败: %v", err)
	}
}
