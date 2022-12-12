package student

import (
	"encoding/json"

	"github.com/OPengXJ/Homework/interner/repository/mysql/student"
)

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (s *Service) Login(req *LoginRequest) ([]byte, error) {
	qb := student.NewQueryBuilder()
	qb.WhereUsername(req.Username)
	qb.WherePassword(req.Password)
	student, err := qb.First(s.db.Read)
	if err != nil {
		return nil, err
	}
	var LoginReponse struct {
		UserId   uint   `json:"userid"`
		UserName string `json:"username"`
	}
	LoginReponse.UserId = student.Id
	LoginReponse.UserName = student.UserName
	byte, err := json.Marshal(LoginReponse)
	if err != nil {
		return nil, err
	}
	return byte, nil
}
