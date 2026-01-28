package master

import (
	"context"
	"fmt"

	"github.com/R2Remote/ChronoGo/internal/infrastructure/redis"
)

func ListenAndDispatch() {
	for {
		cmd := redis.Client.BLPop(context.Background(), 0, "jobs")
		fmt.Println(cmd.Result())
	}
}
