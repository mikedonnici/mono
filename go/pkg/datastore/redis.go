package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisConnection struct {
	conn *redis.Client
}

// newRedisConnection returns a redisConnection value
// dns format: redis://<user>:<pass>@localhost:6379/<db>
func newRedisConnection(dsn string, timeoutSeconds int) (*redisConnection, error) {

	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("could not create redis connection: %w", err)
	}
	cl := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()
	if _, err = cl.Ping(ctx).Result(); err != nil {

		return nil, fmt.Errorf("could not ping redis server: %w", err)
	}
	return &redisConnection{conn: cl}, nil
}
