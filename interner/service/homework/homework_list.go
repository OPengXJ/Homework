package homework

import "github.com/OPengXJ/Homework/interner/repository/mysql/homework"

type SearchHomeworkData struct {
	Page     int      `form:"page"`     //第几页
	PageSize int      `form:"pagesize"` //每页显示条数
	Session  int      `form:"session"`
	College  string   `form:"college"`
	TeaName string `form:"teaname"`
	ClassName string `form:"classname"`
	Order    []string `form:"order"`
}

func (s *Service) HomeworkList(data *SearchHomeworkData) ([]*homework.Homework,error ){
	qb := homework.NewQueryBuilder()
	if data.Session != 0 {
		qb.WhereSession(data.Session)
	}
	if data.College != "" {
		qb.WhereCollege(data.College)
	}
	if data.ClassName !=""{
		qb.WhereClassName(data.ClassName)
	}
	if data.TeaName!=""{
		qb.WhereTeaName(data.TeaName)
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
	HomeworkList,err:=qb.QueryAll(s.db.Read)
	if err!=nil{
		return nil,err
	}
	return HomeworkList,nil
}
