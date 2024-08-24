package component

import (
	"ahyalfan/golang_e_money/domain"
	"context"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

func GetCacheConnection() domain.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatal("error connecting to cache repository: ", err)
	}
	return cache
}
