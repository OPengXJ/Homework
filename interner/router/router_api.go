package router

import (
	"github.com/OPengXJ/GoPro/interner/api/admin"
	"github.com/OPengXJ/Homework/interner/api/class"
	"github.com/OPengXJ/Homework/interner/api/homework"
	"github.com/OPengXJ/Homework/interner/api/student"
	"github.com/OPengXJ/Homework/interner/api/teacher"
	"github.com/OPengXJ/Homework/interner/router/middlewares"
	"github.com/gin-gonic/gin"
)

func setApiRouter(r *gin.Engine){
	
	//每一个部份的处理器结构体
	adminHandler:=admin.New()
	studentHandler:=student.New()
	classHandler:=class.New()
	teacherHandler:=teacher.New()
	homeworkHandler:=homework.New()

	//分类的依据是，当用户登录的是哪一个类别的用户时，它能用到的功能
	//admin
	admin:=r.Group("/api/admin")
	{	
		admin.POST("/login",adminHandler.Login())
		admin.POST("/create",adminHandler.Create())
		adminAuthed:=admin.Group("/")
		adminAuthed.Use(middlewares.JWTAuth("admin"))

		//管理员在非以下功能模块中的功能路径
		{
			adminAuthed.GET("/test",adminHandler.ATest())
		}

		//管理员在老师模块下的功能的路径
		adminTeacher:=adminAuthed.Group("/teacher")
		{
			adminTeacher.POST("/createtea",teacherHandler.Create())
			adminTeacher.GET("/list",teacherHandler.List())
			
		}

		//管理员在学生模块下的功能的路径
		adminStudent:=adminAuthed.Group("/student")
		{
			adminStudent.POST("/createstu",studentHandler.Create())
			adminStudent.GET("/list",studentHandler.List())
		}

		//管理员在班级模块下的功能路径
		adminClass:=adminAuthed.Group("/class")
		{
			adminClass.POST("/create",classHandler.Create())
			adminClass.GET("/list",classHandler.List())
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
		stuClassAuthed:=class.Group("/student")
		{
			stuClassAuthed.GET("/list",classHandler.List())
		}
	}

	//teacher
	//teacher还有作业的功能
	teacher:=r.Group("/api/teacher")
	{
		teacher.POST("/login",teacherHandler.Login())
		teacherAuthed:=teacher.Group("/")
		teacherAuthed.Use(middlewares.JWTAuth("teacher"))
		//老师在作业模块下的功能路径
		teacherHomework:=teacherAuthed.Group("/homework")
		{
			teacherHomework.POST("/create",homeworkHandler.Create())
			teacherHomework.GET("/list",homeworkHandler.ListByES())
			teacherHomework.GET("/detail",homeworkHandler.HomeworkDetail())
		}
		//老师在学生模块下的功能路径
		teacherStudent:=teacherAuthed.Group("/student")
		{
			teacherStudent.GET("/list",studentHandler.List())
		}

	}
	
}