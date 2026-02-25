package strategy

import "github.com/R2Remote/ChronoGo/internal/domain/entity"

type DispatchStrategy interface {
	SelectServers(job *entity.Job) ([]entity.Server, error)
}
