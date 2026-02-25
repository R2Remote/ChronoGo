package strategy

import "github.com/R2Remote/ChronoGo/internal/domain/entity"

type BroadcastStrategy struct {
}

func NewBroadcastStrategy() DispatchStrategy {
	return &BroadcastStrategy{}
}

func (s *BroadcastStrategy) SelectServers(job *entity.Job) ([]entity.Server, error) {
	return nil, nil
}
