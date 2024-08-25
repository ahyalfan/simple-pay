package repository

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/internal/config"
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCacheRepository struct {
	client *redis.Client
}

func NewRedisCache(cnf *config.Config) domain.CacheRepository {
	db, _ := strconv.Atoi(cnf.Redis.DB)

	return &RedisCacheRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     cnf.Redis.Addr,
			Password: cnf.Redis.Password,
			DB:       db,
		}),
	}
}

// Get implements domain.CacheRepository.
func (r *RedisCacheRepository) Get(key string) ([]byte, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, domain.ErrCacheMiss
	} else if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

// Set implements domain.CacheRepository.
func (r *RedisCacheRepository) Set(key string, entry []byte) error {
	return r.client.Set(context.Background(), key, entry, 15*time.Minute).Err() // kita atur 15 menit saja
}
