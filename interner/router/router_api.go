package router

import (
	"github.com/OPengXJ/GoPro/interner/api/admin"
	"github.com/OPengXJ/GoPro/interner/api/student"
	"github.com/OPengXJ/GoPro/interner/router/middlewares"
	"github.com/gin-gonic/gin"
)

func setApiRouter(r *gin.Engine){
	//admin
	adminHandler:=admin.New()
	admin:=r.Group("/api/admin")
	{	
		admin.POST("/login",adminHandler.Login())
		admin.POST("/create",adminHandler.Create())
		adminAuthed:=admin.Group("/")
		adminAuthed.Use(middlewares.JWTAuth("admin"))
		{
			adminAuthed.GET("test",adminHandler.ATest())
		}
	}
	studentHandler:=student.New()
	student:=r.Group("/api/student")
	{
		stuAuthed:=student.Group("/")
		stuAuthed.Use(middlewares.JWTAuth("student"))
		{
			stuAuthed.GET("test",studentHandler.ATest())
		}
	}

}