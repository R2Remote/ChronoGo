package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/redis/go-redis/v9"
)

const queueName = "task_queue"

type TaskConsumer struct {
	dispatcher  *Dispatcher
	redisClient *redis.Client
	queueName   string
	ctx         context.Context
}

func NewTaskConsumer(dispatcher *Dispatcher, redisClient *redis.Client, ctx context.Context) *TaskConsumer {
	return &TaskConsumer{dispatcher: dispatcher, redisClient: redisClient, queueName: queueName, ctx: ctx}
}

func (c *TaskConsumer) FetchTask() *entity.Task {
	// BLPop returns []string{key, value}
	result, err := c.redisClient.BLPop(c.ctx, 0, c.queueName).Result()
	if err != nil || len(result) < 2 {
		log.Println("Consumer: get task error:", err)
	}

	val := result[1]
	var task entity.Task
	if err := json.Unmarshal([]byte(val), &task); err != nil {
		log.Printf("Consumer: invalid task json: %s, error: %v", val, err)
	}
	return &task
}
