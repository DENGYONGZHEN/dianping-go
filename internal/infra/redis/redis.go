package redis

import (
	"context"
	"hm-dianping-go/internal/constant"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func New(addr string) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{client: rdb}, nil
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

// 向redis设置验证码
func (r *RedisClient) SetLoginCode(ctx context.Context, phone string, code string) error {

	key := constant.LOGIN_CODE_KEY + phone

	ttl := time.Duration(constant.LOGIN_CODE_TTL) * time.Minute

	return r.client.Set(ctx, key, code, ttl).Err()
}
