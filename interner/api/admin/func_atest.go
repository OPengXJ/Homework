package admin

import (
	"net/http"

	"github.com/OPengXJ/GoPro/interner/repository/redis"

	"github.com/gin-gonic/gin"
)

func (h *Handle) ATest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cache := redis.GetRedisCache()
		cache.Redis.Set(ctx,"hello", "world", 0)
		ctx.String(http.StatusOK, cache.Redis.Get(ctx,"hello").String())
	}
}
