package configs

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func ClientRedis() *redis.Client {
	addr := fmt.Sprintf("%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
