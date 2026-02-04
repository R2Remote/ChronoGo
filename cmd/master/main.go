package main

import (
	"fmt"
	"log"

	"github.com/R2Remote/ChronoGo/internal/config"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/database"
	"github.com/R2Remote/ChronoGo/internal/infrastructure/redis"
	"github.com/R2Remote/ChronoGo/internal/interfaces/api"
	"github.com/R2Remote/ChronoGo/internal/interfaces/master"
)

func main() {
	//load config
	if err := config.LoadConfig(""); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	//init database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	//init redis
	if err := redis.InitRedis(); err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}
	defer redis.CloseRedis()

	//listen and dispatch
	// 使用 Init 进行初始化，内部启动消费者
	master.Init()
	fmt.Println("🚀 Master 节点启动成功！正在监听任务队列...")

	// 启动 Web API 服务 (对接 FlowBoard)
	apiServer := api.NewServer()
	if err := apiServer.Start(":8082"); err != nil {
		log.Fatalf("Master Web API 启动失败: %v", err)
	}
	fmt.Println("🚀 Master Web API 启动成功！")
}
