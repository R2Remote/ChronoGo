package worker

import (
	context "context"
	"fmt"
	"log"
	"net"

	pb "github.com/R2Remote/ChronoGo/api/proto"

	"google.golang.org/grpc"
)

type JobHandler func(params string) error

type Server struct {
	pb.UnimplementedJobWorkerServer
	handlers map[string]JobHandler // 存储 任务名 -> 函数 的映射
}

// 供业务方注册函数
func (s *Server) RegisterHandler(name string, handler JobHandler) {
	s.handlers[name] = handler
}

// ExecuteJob 执行任务
func (s *Server) ExecuteJob(ctx context.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	log.Printf("收到任务执行请求: ID=%d, Handler=%s, Params=%s", req.JobId, req.JobHandler, req.JobParams)
	handler, ok := s.handlers[req.JobHandler]
	if !ok {
		return &pb.ExecuteResponse{Success: false, Msg: "未找到该处理器"}, nil
	}

	// 运行业务逻辑
	err := handler(req.JobParams)
	if err != nil {
		return &pb.ExecuteResponse{Success: false, Msg: err.Error()}, nil
	}

	return &pb.ExecuteResponse{Success: true}, nil
}

// NewWorker 创建并初始化 SDK 实例
func NewWorker() *Server {
	return &Server{
		handlers: make(map[string]JobHandler),
	}
}

// ChronoGo/sdk/worker/server.go

// Start 启动 gRPC 服务并阻塞（业务方通常会 go srv.Start()）
func (s *Server) Start(port int) error {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterJobWorkerServer(grpcServer, s) // 注意这里传的是 s，即 *Server

	log.Printf("Worker gRPC Server started on %s", address)
	return grpcServer.Serve(listener)
}
