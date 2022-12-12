package class

import (
	"strconv"

	"github.com/OPengXJ/Homework/interner/repository/mysql/class"
)

type CreateClassData struct{
	Classname string `form:"classname" binding:"required"`
	Session string `form:"session" binding:"required"`
	College string `form:"college" binding:"required"`
}
func(s *Service)Create(data *CreateClassData)error{
	model:=class.NewModel()
	var parseErr error
	model.Classname=data.Classname
	model.College=data.College
	model.Session,parseErr=strconv.Atoi(data.Session)
	if parseErr!=nil{
		return nil
	}
	err:=model.Create(s.db.Write)
	if err!=nil{
		return err
	}
	return nil
}
