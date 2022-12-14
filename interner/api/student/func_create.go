package student

import (
	"fmt"
	"log"

	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/GoPro/interner/service/student"
	"github.com/gin-gonic/gin"
)

func (h *Handle) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createData := &student.CreateStudentData{}
		if err := ctx.ShouldBind(createData); err != nil {
			log.Println(createData)
		} else {
			log.Println(err)
		}
		fmt.Println(createData)
		repo := mysql.GetMysqlRepo()
		service := student.New(*repo,ctx)
		if err := service.Create(createData); err != nil {
			log.Println("Create failed")
			return
		}
		ctx.String(200, "εε»Ίζε")
	}
}
