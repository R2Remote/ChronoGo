package main

import (
	"fmt"
	"log"

	"ChronoGo/internal/common"
	"ChronoGo/internal/interfaces/api"
)

func main() {
	// 1. 初始化 Etcd 连接
	cli, err := common.NewEtcdClient([]string{"localhost:2379"})
	if err != nil {
		log.Fatalf("Master 启动失败: 无法连接 Etcd, %v", err)
	}
	defer cli.Close()

	fmt.Println("🚀 Master 节点启动成功！正在监听任务队列...")

	// 启动 Web API 服务 (对接 FlowBoard)
	apiServer := api.NewServer()
	if err := apiServer.Start(":8080"); err != nil {
		log.Fatalf("Master Web API 启动失败: %v", err)
	}
}
