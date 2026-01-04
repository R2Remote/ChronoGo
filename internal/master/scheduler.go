package master

// Scheduler 负责 Master 的调度逻辑
// 功能：Master 随便挑一个活着的 Worker，调用它的 RunJob
//
// TODO: 实现以下功能
// - 写一个死循环，每隔 5 秒，从内存 Map 里拿一个 Worker IP
// - 建立 gRPC 连接，调用 RunJob
// - 看到 Worker 控制台打印 "收到任务，正在执行..."

