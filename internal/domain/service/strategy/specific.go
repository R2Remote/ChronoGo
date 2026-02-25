package strategy

import "github.com/R2Remote/ChronoGo/internal/domain/entity"

type SpecificStrategy struct {
}

func NewSpecificStrategy() DispatchStrategy {
	return &SpecificStrategy{}
}

func (s *SpecificStrategy) SelectServers(job *entity.Job) ([]entity.Server, error) {
	return nil, nil
}
