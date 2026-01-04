package master

// ServiceDiscovery 负责 Master 的服务发现逻辑
// 功能：实时感知 Worker 的上线和下线，不需要重启 Master
// 技术点：Etcd Watch 机制
//
// 逻辑：
// 1. 先 Get 前缀 /honor-cron/workers/ 下的所有现有节点，存到内存 Map 里
// 2. 启动 Watch 协程，监听这个前缀的 PUT (上线) 和 DELETE (下线) 事件，实时更新内存 Map
//
// TODO: 实现以下功能
// - 启动 Master，启动 Worker，Master 控制台能打印 "检测到新节点加入"
// - 关掉 Worker，Master 能打印 "检测到节点下线"

