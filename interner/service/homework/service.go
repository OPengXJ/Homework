package homework

import (
	"context"

	"github.com/OPengXJ/Homework/interner/repository/mysql"
	myredis "github.com/OPengXJ/Homework/interner/repository/redis"
)

type Service struct {
	db    mysql.Repo
	cache *myredis.Cache
	ctx   context.Context
}

func New(db mysql.Repo, ctx context.Context) *Service {
	return &Service{
		db:    db,
		cache: myredis.GetRedisCache(),
		ctx:   ctx,
	}
}
