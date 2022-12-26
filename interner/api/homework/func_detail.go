package homework

import (
	"net/http"

	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/service/homework"
	"github.com/gin-gonic/gin"
)

func(h *Handle)HomeworkDetail()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		detailSearchData:=&homework.HomeworkDetailSearchData{}
		err:=ctx.ShouldBind(detailSearchData)
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,err.Error())
		}
		rep:=mysql.GetMysqlRepo()
		service:=homework.New(rep,ctx)
		result,err:=service.HomeworkDeatil(detailSearchData)
		if err!=nil{
			ctx.String(http.StatusBadRequest,err.Error())
		}
		ctx.String(http.StatusOK,result)
	}
}