package admin

import (
	"context"

	"github.com/OPengXJ/GoPro/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/repository/redis"
)


type Service struct{
	db mysql.Repo
	cache redis.Cache
	ctx context.Context
}

func New(db mysql.Repo,ctx context.Context)*Service{
	return &Service{
		db: db,
		cache: *redis.GetRedisCache(),
		ctx: ctx,
	}
}