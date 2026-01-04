package worker

// Register 负责 Worker 的服务注册逻辑
// 功能：启动时把自己 IP 写进 Etcd，挂掉时自动消失
// 技术点：Etcd Grant (租约) + KeepAlive (自动续期)
// Key 格式：/honor-cron/workers/{IP}
//
// TODO: 实现以下功能
// - 启动时创建 Lease
// - 开启 KeepAlive 协程
// - 写入 Etcd，模拟 kill 掉进程，看 Key 是否自动消失

