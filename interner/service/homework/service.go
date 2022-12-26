package homework

import (
	"context"

	"github.com/OPengXJ/Homework/interner/repository/elasticsearch"
	"github.com/OPengXJ/Homework/interner/repository/mysql"
	myredis "github.com/OPengXJ/Homework/interner/repository/redis"
	"github.com/olivere/elastic/v7"
)

type Service struct {
	db    *mysql.Repo
	cache *myredis.Cache
	es	*elastic.Client
	ctx   context.Context
}

func New(db *mysql.Repo, ctx context.Context) *Service {
	return &Service{
		db:    db,
		cache: myredis.GetRedisCache(),
		es:elasticsearch.GetEsClient(),
		ctx:   ctx,
	}
}
