package service

import (
	"context"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/redis/go-redis/v9"
)

type ServerService struct {
	redisClient *redis.Client
}

func (s *ServerService) Loop(m *map[uint64]entity.Server) {

}

func (s *ServerService) List(ctx context.Context, page int, pageSize int) ([]*entity.Server, int64) {
	panic("unimplemented")
}

func NewServerService(redisClient *redis.Client) *ServerService {
	return &ServerService{redisClient: redisClient}
}
