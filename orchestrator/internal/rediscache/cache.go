package rediscache

import (
	"context"
	"time"
)

func (s *Service) Get(ctx context.Context, key string) (string, error) {
	return s.redisClient.Get(ctx, key).Result()
}

func (s *Service) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.redisClient.Set(ctx, key, value, expiration).Err()
}

func (s *Service) Del(ctx context.Context, keys ...string) error {
	return s.redisClient.Del(ctx, keys...).Err()
}

func (s *Service) Keys(ctx context.Context, pattern string) ([]string, error) {
	return s.redisClient.Keys(ctx, pattern).Result()
}
