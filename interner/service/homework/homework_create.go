package homework

import (
	"strconv"
	"time"

	"github.com/OPengXJ/Homework/interner/repository/mysql/homework"
)

type CreateHomeworkData struct{
	College   string `form:"college"`
	TeaName string	`form:"teaname"`
	ClassName string	`form:"classname"`
	Content   string	`form:"content"`
	StartTime string	`form:"starttime"`
	Deadline  string	`form:"deadline"`
	WorkName  string	`form:"workname"`
	Session string `form:"session"`
}
func(s *Service)Create(data *CreateHomeworkData)error{
	model:=homework.NewModel()
	var parseErr error
	model.College=data.College
	model.TeaName=data.TeaName
	model.ClassName=data.ClassName
	model.Content=data.Content
	model.WorkName=data.WorkName
	//这里假设从前端传入的格式为 "2022-12-13T10:11"
	timeFormat:="2006-01-02T15:04"
	startTime,err:=time.ParseInLocation(timeFormat,data.StartTime,time.Local)
	if err!=nil{
		return err
	}
	deadLine,err:=time.ParseInLocation(timeFormat,data.Deadline,time.Local)
	if err!=nil{
		return err
	}
	model.StartTime=startTime
	model.Deadline=deadLine
	model.Session,parseErr=strconv.Atoi(data.Session)
	if parseErr!=nil{
		return nil
	}
	err=model.Create(s.db.Write)
	if err!=nil{
		return err
	}
	return nil
}
