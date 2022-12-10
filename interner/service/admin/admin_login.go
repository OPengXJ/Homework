package admin

import (
	"encoding/json"

	"github.com/OPengXJ/GoPro/interner/repository/mysql/admin"
)

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}


func (s *Service) Login(req *LoginRequest) ([]byte, error) {
	qb := admin.NewQueryBuilder()
	qb.WhereUsername(req.Username)
	qb.WherePassword(req.Password)
	admin, err := qb.First(s.db.Read)
	if err != nil {
		return nil, err
	}
	var LoginReponse struct{
		UserId uint `json:"userid"`
		UserName string `json:"username"`
	}
	LoginReponse.UserId=admin.ID
	LoginReponse.UserName=admin.UserName
	byte,err:=json.Marshal(LoginReponse)
	if err!=nil{
		return nil,err
	}
	return byte, nil
}
