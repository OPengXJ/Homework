package student

import (
	"net/http"

	"github.com/OPengXJ/Homework/interner/pkg/token"
	"github.com/OPengXJ/Homework/interner/repository/mysql"
	"github.com/OPengXJ/Homework/interner/service/student"
	"github.com/gin-gonic/gin"
)


func(h *Handle)Login()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		req:=&student.LoginRequest{}
		err:=ctx.ShouldBind(&req)
		if err!=nil{
			ctx.String(http.StatusBadRequest,"绑定信息出错：%v",err)
			ctx.Abort()
			return
		}
		rep:=mysql.GetMysqlRepo()
		service:=student.New(*rep,ctx)
		res,err:=service.Login(req)
		if err!=nil{
			ctx.String(http.StatusBadRequest,"登陆名或者密码不正确：%v",err)
			ctx.Abort()
			return
		}
		tokenString:=token.CreateToken(ctx,res,"student")
		ctx.SetCookie("token",tokenString,3600*24,"/","",false,true)
		ctx.Header("x-token",tokenString)
		ctx.JSON(http.StatusOK,gin.H{
			"user":res,
			"token":tokenString,
		})
	}
}