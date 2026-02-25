package service

import (
	"context"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/redis/go-redis/v9"
)

type ServerService struct {
	redisClient *redis.Client
}

func (s *ServerService) List(ctx context.Context, page int, pageSize int) ([]*entity.Server, int64) {
	// return s.redisClient.Keys(ctx, "server:*")
	return nil, 0
}

func NewServerService(redisClient *redis.Client) *ServerService {
	return &ServerService{redisClient: redisClient}
}
