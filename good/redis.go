package good

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// RedisSource interface is used to do a DI on GetKey and SetKey functions (accept
// interfaces return struct).
type RedisSource interface {
	Get(context.Context, string) *redis.StringCmd
	Set(context.Context, string, interface{}, time.Duration) *redis.StatusCmd
}

type MessageBroker struct {
	*redis.Client
}

func NewClient(connectionString string) (MessageBroker, error) {
	opt, err := redis.ParseURL(connectionString)
	return MessageBroker{redis.NewClient(opt)}, err
}

func GetKey(rdb RedisSource, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	return val, err
}

func SetKey(rdb RedisSource, key, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}
