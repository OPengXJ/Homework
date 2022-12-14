package homework

import (
	"encoding/json"
	"net/http"

	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/service/homework"
	"github.com/gin-gonic/gin"
)

func (h *Handle) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchHomeworkData := &homework.SearchHomeworkData{}
		err := ctx.ShouldBind(searchHomeworkData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		rep := mysql.GetMysqlRepo()
		service := homework.New(*rep)
		HomeworkList, err := service.HomeworkList(searchHomeworkData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		byteData, err := json.Marshal(HomeworkList)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK,string(byteData))
	}
}
