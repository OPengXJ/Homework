package student

import (
	"github.com/OPengXJ/Homework/interner/repository/mysql/student"
)

type CreateStudentData struct {
	Username string `form:"username" bind:"required"`
	Password string `form:"password" bind:"required"`
	StuClass string `form:"stuclass" bind:"required"`
	StuMajor string `form:"stumajor" bind:"required"`
	StuName  string `form:"stuname" bind:"required"`
}

func (s *Service) Create(data *CreateStudentData) error {
	model := student.NewModel()
	model.UserName = data.Username
	model.UserPass = data.Password
	model.StuClass = data.StuClass
	model.StuMajor = data.StuMajor
	model.StuName = data.StuName
	err := model.Create(s.db.Write)
	if err != nil {
		return err
	}
	return nil
}
