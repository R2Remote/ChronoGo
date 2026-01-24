package master

import (
	context "context"
	"fmt"
	"log"
	"time"

	pb "ChronoGo/api/proto"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ExecuteJobClient 模拟向指定 Worker 下发任务
func ExecuteJobClient(addr string) error {
	// 1. 建立 gRPC 连接
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("连接 Worker %s 失败: %v", addr, err)
	}
	defer conn.Close()

	client := pb.NewJobWorkerClient(conn)

	// 2. 构造请求参数
	req := &pb.ExecuteRequest{
		JobId:      time.Now().Unix(),
		JobHandler: "some_task",
		JobParams:  "key=value",
	}

	// 3. 发送请求 (设置 5 秒超时)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.ExecuteJob(ctx, req)
	if err != nil {
		return fmt.Errorf("调用 ExecuteJob 失败: %v", err)
	}

	log.Printf("Worker 响应: [Success: %v] [Msg: %s]", resp.Success, resp.Msg)
	return nil
}

// StartScheduler 模拟 Master 的调度。第一阶段只是死循环给同一个 Worker 发请求
func StartScheduler() {
	log.Println("⏰ Master 调度循环开始...")
	target := "localhost:8888"

	for {
		log.Printf("⏳ 正在尝试给 Worker (%s) 下派任务...", target)
		if err := ExecuteJobClient(target); err != nil {
			log.Printf("❌ 任务下派失败: %v", err)
		}

		// 每 10 秒钟跑一个模拟任务
		time.Sleep(10 * time.Second)
	}
}
