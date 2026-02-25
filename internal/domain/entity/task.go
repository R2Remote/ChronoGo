package entity

type Task struct {
	JobID    uint64
	Param    string
	Strategy StrategyType
	ServerID uint64
}

type StrategyType int

const (
	Random StrategyType = iota // 使用 iota 自动递增更优雅
	RoundRobin
	Broadcast
	Specific
)
