package redis

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	goredis "github.com/redis/go-redis/v9"
)

func NewRedisClient(redisURL string) (*goredis.Client, error) {
	u, err := url.Parse(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}

	password, _ := u.User.Password()

	addr := u.Host
	if u.Port() == "" {
		addr = u.Host + ":6379"
	}

	db := 0
	if u.Path != "" && u.Path != "/" {
		fmt.Sscanf(u.Path, "/%d", &db)
	}

	rdb := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	slog.Info("connected to Redis", "addr", addr)
	return rdb, nil
}
