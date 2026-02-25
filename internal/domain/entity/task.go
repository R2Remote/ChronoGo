package entity

type Task struct {
	JobID    uint64
	Param    string
	Strategy ExecuteStrategy
	ServerID uint64
}

type ExecuteStrategy int

const (
	// 随机
	ExecuteStrategy_Random ExecuteStrategy = 0
	// 轮询
	ExecuteStrategy_RoundRobin ExecuteStrategy = 1
	// 广播
	ExecuteStrategy_Broadcast ExecuteStrategy = 2
	// 指定server
	ExecuteStrategy_Specific ExecuteStrategy = 3
)
