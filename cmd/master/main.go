package main

import (
	"fmt"

	"ChronoGo/internal/master"
)

func main() {

	fmt.Println("🚀 Master 节点启动成功！正在监听任务队列...")

	// 启动调度循环 (模拟)
	go master.StartScheduler()

	// TODO: 启动服务发现
	// master.StartServiceDiscovery(cli)

	// 阻塞主进程
	select {}
}
