package student

import (
	"encoding/json"
	"net/http"

	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/service/student"
	"github.com/gin-gonic/gin"
)

func (h *Handle) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchStudentData := &student.SearchStudentData{}
		err := ctx.ShouldBind(searchStudentData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		rep := mysql.GetMysqlRepo()
		service := student.New(*rep)
		StudentList, err := service.StudentList(searchStudentData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		byteData, err := json.Marshal(StudentList)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK,string(byteData))
	}
}
