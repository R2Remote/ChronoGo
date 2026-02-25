package strategy

import "github.com/R2Remote/ChronoGo/internal/domain/entity"

type RandomStrategy struct {
}

func NewRandomStrategy() DispatchStrategy {
	return &RandomStrategy{}
}

func (s *RandomStrategy) SelectServers(job *entity.Job) ([]entity.Server, error) {
	return nil, nil
}
