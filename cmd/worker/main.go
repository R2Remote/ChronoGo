package main

import (
	"fmt"
	"log"

	"go-chrono/internal/common"
	// TODO: 导入 worker 包
	// "go-chrono/internal/worker"
)

func main() {
	// Worker 也要连 Etcd，用来注册自己
	cli, err := common.NewEtcdClient([]string{"localhost:2379"})
	if err != nil {
		log.Fatalf("Worker 启动失败: %v", err)
	}
	defer cli.Close()

	fmt.Println("👷 Worker 节点已上线！等待 Master 派活...")

	// TODO: 启动服务注册（保活）
	// worker.Register(cli)

	// TODO: 启动 gRPC Server
	// worker.StartGRPCServer()

	// 阻塞主进程
	select {}
}