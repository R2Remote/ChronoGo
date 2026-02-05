package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	domainService "github.com/R2Remote/ChronoGo/internal/domain/service"
	"github.com/redis/go-redis/v9"
)

type TaskConsumer struct {
	dispatcher  *domainService.Dispatcher
	redisClient *redis.Client
	queueName   string
}

func NewTaskConsumer(dispatcher *domainService.Dispatcher, redisClient *redis.Client) *TaskConsumer {
	return &TaskConsumer{
		dispatcher:  dispatcher,
		redisClient: redisClient,
		queueName:   "job_queue",
	}
}

func (c *TaskConsumer) Start() {
	go c.loop()
}

func (c *TaskConsumer) loop() {
	for {
		// BLPop returns []string{key, value}
		result, err := c.redisClient.BLPop(context.Background(), 0, c.queueName).Result()
		if err != nil {
			log.Println("Consumer: get task error:", err)
			continue
		}
		if len(result) < 2 {
			continue
		}

		val := result[1]
		var task entity.Task
		if err := json.Unmarshal([]byte(val), &task); err != nil {
			log.Printf("Consumer: invalid task json: %s, error: %v", val, err)
			continue
		}

		// Delegate to dispatcher
		c.dispatcher.Dispatch(&task)
	}
}
