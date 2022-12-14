package teacher

import (
	"encoding/json"
	"net/http"

	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/service/teacher"
	"github.com/gin-gonic/gin"
)

func (h *Handle) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchTeacherData := &teacher.SearchTeacherData{}
		err := ctx.ShouldBind(searchTeacherData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		rep := mysql.GetMysqlRepo()
		service := teacher.New(*rep)
		TeacherList, err := service.TeacherList(searchTeacherData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		byteData, err := json.Marshal(TeacherList)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK,string(byteData))
	}
}
