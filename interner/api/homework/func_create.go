package homework

import (
	"fmt"
	"log"
	"net/http"
	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/service/homework"
	"github.com/gin-gonic/gin"
)

func (h *Handle) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createData := &homework.CreateHomeworkData{}
		if err := ctx.ShouldBind(createData); err != nil {
			log.Println(createData)
		} else {
			log.Println(err)
		}
		fmt.Println(createData)
		repo := mysql.GetMysqlRepo()
		service := homework.New(*repo)
		if err := service.Create(createData); err != nil {
			log.Println("Create failed")
			ctx.AbortWithError(http.StatusBadRequest,err)
			return
		}
		ctx.String(200, "创建成功")
	}
}