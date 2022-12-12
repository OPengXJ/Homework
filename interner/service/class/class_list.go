package class

import "github.com/OPengXJ/Homework/interner/repository/mysql/class"

type SearchCLassData struct {
	Page     int      `form:"page"`     //第几页
	PageSize int      `form:"pagesize"` //每页显示条数
	Session  int      `form:"session"`
	College  string   `form:"college"`
	Order    []string `form:"order"`
}

func (s *Service) ClassList(data *SearchCLassData) ([]*class.Class,error ){
	qb := class.NewQueryBuilder()
	if data.Session != 0 {
		qb.WhereSession(data.Session)
	}
	if data.College != "" {
		qb.WhereCollege(data.College)
	}
	page := data.Page
	if page == 0 {
		page = 1
	}
	pageSize := data.PageSize
	if pageSize == 0 {
		pageSize = 5
	}
	qb.Limit(pageSize)
	offset := (page - 1) * pageSize
	qb.OffSet(offset)
	qb.Order(data.Order)
	ClassList,err:=qb.QueryAll(s.db.Read)
	if err!=nil{
		return nil,err
	}
	return ClassList,nil
}
