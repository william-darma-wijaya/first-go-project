package caching

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache struct {
	client *redis.Client
}

func NewUserCache(client *redis.Client) *UserCache {
	return &UserCache{client: client}
}

func userKey(id string) string {
	return "user:" + id
}

func (c *UserCache) Get(ctx context.Context, token string) (string, error) {
	return c.client.Get(ctx, "auth:"+token).Result()
}

func (c *UserCache) Set(ctx context.Context, token string, userID string, ttl int) error {
	return c.client.Set(ctx, "auth:"+token, userID, time.Duration(ttl)*time.Minute).Err()
}

func (c *UserCache) Delete(ctx context.Context, token string) error {
	return c.client.Del(ctx, "auth:"+token).Err()
}
