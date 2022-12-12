package router

import (
	"github.com/OPengXJ/GoPro/interner/api/admin"
	"github.com/OPengXJ/Homework/interner/api/class"
	"github.com/OPengXJ/Homework/interner/api/student"
	"github.com/OPengXJ/Homework/interner/router/middlewares"
	"github.com/gin-gonic/gin"
)

func setApiRouter(r *gin.Engine){
	
	//每一个部份的处理器结构体
	adminHandler:=admin.New()
	studentHandler:=student.New()
	classHandler:=class.New()

	//admin
	admin:=r.Group("/api/admin")
	{	
		admin.POST("/login",adminHandler.Login())
		admin.POST("/create",adminHandler.Create())
		adminAuthed:=admin.Group("/")
		adminAuthed.Use(middlewares.JWTAuth("admin"))
		{
			adminAuthed.POST("/createstu",studentHandler.Create())
			adminAuthed.GET("test",adminHandler.ATest())
		}
	}

	//student
	student:=r.Group("/api/student")
	{
		student.POST("/login",studentHandler.Login())
		stuAuthed:=student.Group("/")
		stuAuthed.Use(middlewares.JWTAuth("student"))
		{
			stuAuthed.GET("test",studentHandler.ATest())
		}
	}

	//class
	class:=r.Group("/api/class")
	{
		adminClassAuthed:=class.Group("/admin")
		adminClassAuthed.Use(middlewares.JWTAuth("admin"))
		{
			adminClassAuthed.GET("/list",classHandler.List())
			adminClassAuthed.POST("/create",classHandler.Create())
		}
		stuClassAuthed:=class.Group("/student")
		{
			stuClassAuthed.GET("/list",classHandler.List())
		}
	}


}