package redis

import (
	"context"

	"github.com/OPengXJ/GoPro/configs"
	"github.com/alibaba/tair-go/tair"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Redis     *redis.Client
	TairRedis *tair.TairClient
}

var cache = new(Cache)

func init() {
	var err error
	cache, err = New()
	if err != nil {
		panic(err)
	}
}

func New() (*Cache, error) {
	cfg := configs.Get().Redis
	ctx := context.Background()
	redisclient := redis.NewClient(&redis.Options{
		Addr:         cfg.DBaddr,
		Password:     cfg.DBpass,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	if err := redisclient.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	TairClient := tair.NewTairClient(&redis.Options{
		Addr:         cfg.DBaddr,
		Password:     cfg.DBpass,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	if err := TairClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Cache{
		Redis:     redisclient,
		TairRedis: TairClient,
	}, nil
}
func GetRedisCache() *Cache {
	return cache
}
