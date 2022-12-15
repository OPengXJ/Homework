package class

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/service/class"
	"github.com/gin-gonic/gin"
)

func (h *Handle) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createData := &class.CreateClassData{}
		if err := ctx.ShouldBind(createData); err != nil {
			log.Println(createData)
		} else {
			log.Println(err)
		}
		fmt.Println(createData)
		repo := mysql.GetMysqlRepo()
		service := class.New(*repo,ctx)
		if err := service.Create(createData); err != nil {
			log.Println("Create failed")
			ctx.AbortWithError(http.StatusBadRequest,err)
			return
		}
		ctx.String(200, "创建成功")
	}
}
