package strategy

import "github.com/R2Remote/ChronoGo/internal/domain/entity"

type RoundRobinStrategy struct {
}

func NewRoundRobinStrategy() DispatchStrategy {
	return &RoundRobinStrategy{}
}

func (s *RoundRobinStrategy) SelectServers(job *entity.Job) ([]entity.Server, error) {
	return nil, nil
}
