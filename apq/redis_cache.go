package apq

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// Cache Struct
type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

// NewCache returns a new redis cache client
func NewCache(ttl time.Duration) (*Cache, error) {

	// redisAddress : Redis Host Address
	redisAddress := os.Getenv("REDIS_ADDRESS")

	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Cache{client: client, ttl: ttl}, nil
}

// Add adds data to the cache
func (c *Cache) Add(ctx context.Context, hash string, query string) {
	c.client.Set(apqPrefix+hash, query, c.ttl)
}

// Get retrieves data from the cache
func (c *Cache) Get(ctx context.Context, hash string) (string, bool) {
	s, err := c.client.Get(apqPrefix + hash).Result()
	if err != nil {
		return "", false
	}
	return s, true
}
