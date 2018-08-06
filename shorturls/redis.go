package shorturls

import (
	"encoding/hex"
	"math/rand"

	"github.com/go-redis/redis"
)

// RedisStore is the adapter for storing shorturls in redis
type RedisStore struct {
	client *redis.Client
}

// NewRedisStore returns a new RedisStore
func NewRedisStore(opts *redis.Options) *RedisStore {
	client := redis.NewClient(opts)
	return &RedisStore{client}
}

// Save saves a shorturl in redis
func (s *RedisStore) Save(target string) (string, error) {
	var id string

	for {
		id = randomID()
		_, err := s.client.Get(id).Result()
		if err == redis.Nil {
			break
		}
	}

	err := s.client.Set(id, target, 0).Err()
	if err != nil {
		return "", err
	}

	return id, nil
}

// Get gets a shorturl from redis
func (s *RedisStore) Get(id string) (string, error) {
	return s.client.Get(id).Result()
}

func randomID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return hex.EncodeToString(b)
}
