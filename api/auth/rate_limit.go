package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func IsRateLimited(rdb *redis.Client, phone string) (bool, error) {
	key := fmt.Sprintf("rl:%s", phone)
	count, err := rdb.Incr(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	if count == 1 {
		rdb.Expire(context.Background(), key, 1*time.Minute)
	}
	return count > 5, nil // не более 5 попыток в минуту
}
