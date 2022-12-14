package student

import "github.com/OPengXJ/Homework/interner/repository/mysql/student"

type SearchStudentData struct {
	Page       int      `form:"page"`     //第几页
	PageSize   int      `form:"pagesize"` //每页显示条数
	StuSession int      `form:"stusession"`
	StuCollege string   `form:"stucollege"`
	StuMajor   string   `form:"stumajor"`
	StuClass   string   `form:"stuclass"`
	Order      []string `form:"order"`
}

func (s *Service) StudentList(data *SearchStudentData) ([]*student.Student, error) {
	qb := student.NewQueryBuilder()
	if data.StuSession != 0 {
		qb.WhereStuSession(data.StuSession)
	}
	if data.StuCollege != "" {
		qb.WhereStuCollege(data.StuCollege)
	}
	if data.StuClass != "" {
		qb.WhereStuClass(data.StuClass)
	}
	if data.StuMajor != "" {
		qb.WhereStuMajor(data.StuMajor)
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
	StudentList, err := qb.QueryAll(s.db.Read)
	if err != nil {
		return nil, err
	}
	return StudentList, nil
}
