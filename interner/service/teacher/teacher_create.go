package teacher

import (
	"github.com/OPengXJ/Homework/interner/repository/mysql/teacher"
)

type CreateTeacherData struct {
	Username string `form:"username" bind:"required"`
	Password string `form:"password" bind:"required"`
	TeaName  string `form:"teaname" bind:"required"`
	TeaCollege string `form:"teacollege" bind:"required"`
}

func (s *Service) Create(data *CreateTeacherData) error {
	model := teacher.NewModel()
	model.UserName = data.Username
	model.UserPass = data.Password
	model.TeaName = data.TeaName
	model.TeaCollege= data.TeaCollege
	err := model.Create(s.db.Write)
	if err != nil {
		return err
	}
	return nil
}
