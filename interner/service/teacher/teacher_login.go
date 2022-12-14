package teacher

import (
	"encoding/json"

	"github.com/OPengXJ/Homework/interner/repository/mysql/teacher"
)

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (s *Service) Login(req *LoginRequest) ([]byte, error) {
	qb := teacher.NewQueryBuilder()
	qb.WhereUsername(req.Username)
	qb.WherePassword(req.Password)
	teacher, err := qb.First(s.db.Read)
	if err != nil {
		return nil, err
	}
	var LoginReponse struct {
		UserId   uint   `json:"userid"`
		UserName string `json:"username"`
	}
	LoginReponse.UserId = teacher.Id
	LoginReponse.UserName = teacher.UserName
	byte, err := json.Marshal(LoginReponse)
	if err != nil {
		return nil, err
	}
	return byte, nil
}
