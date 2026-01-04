package worker

// StartGRPCServer 启动 gRPC Server
// 功能：打开一个端口（比如 8081），等待 Master 连上来
// 技术点：实现 WorkerService 接口，在 RunJob 方法里打印日志
//
// TODO: 实现以下功能
// - 监听 TCP 端口
// - 注册 gRPC 服务
// - 阻塞主进程 (select{})

