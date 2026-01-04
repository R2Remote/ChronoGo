package main

import (
	"fmt"
	"log"

	"go-chrono/internal/common"
	// TODO: 导入 master 包
	// "go-chrono/internal/master"
)

func main() {
	// 1. 初始化 Etcd 连接
	// 真实项目中，这里的地址应该从配置文件读取 (Viper)
	cli, err := common.NewEtcdClient([]string{"localhost:2379"})
	if err != nil {
		log.Fatalf("Master 启动失败: 无法连接 Etcd, %v", err)
	}
	defer cli.Close()

	fmt.Println("🚀 Master 节点启动成功！正在监听任务队列...")

	// TODO: 启动服务发现
	// master.StartServiceDiscovery(cli)

	// TODO: 启动调度器
	// master.StartScheduler(cli)

	// 阻塞主进程
	select {}
}