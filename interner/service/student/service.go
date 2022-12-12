package student

import (
	"github.com/OPengXJ/GoPro/interner/repository/mysql"
	"github.com/go-redis/redis"
	myredis "github.com/OPengXJ/Homework/interner/repository/redis"
)


type Service struct{
	db mysql.Repo
	redis *redis.Client

}

func New(db mysql.Repo)*Service{
	return &Service{
		db: db,
		redis: myredis.GetRedisRepo(),
	}
}