package class

import (
	"encoding/json"
	"net/http"

	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/service/class"
	"github.com/gin-gonic/gin"
)

func (h *Handle) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchClassData := &class.SearchCLassData{}
		err := ctx.ShouldBind(searchClassData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		rep := mysql.GetMysqlRepo()
		service := class.New(*rep)
		ClassList, err := service.ClassList(searchClassData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		byteData, err := json.Marshal(ClassList)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK,string(byteData))
	}
}
