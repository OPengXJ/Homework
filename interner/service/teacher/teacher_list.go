package teacher

import "github.com/OPengXJ/Homework/interner/repository/mysql/teacher"

type SearchTeacherData struct {
	Page       int      `form:"page"`     //第几页
	PageSize   int      `form:"pagesize"` //每页显示条数
	Order      []string `form:"order"`
	TeaCollege string   `form:"teacollege"`
}

func (s *Service) TeacherList(data *SearchTeacherData) ([]*teacher.Teacher, error) {
	qb := teacher.NewQueryBuilder()
	if data.TeaCollege != "" {
		qb.WhereTeaCollege(data.TeaCollege)
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
	TeacherList, err := qb.QueryAll(s.db.Read)
	if err != nil {
		return nil, err
	}
	return TeacherList, nil
}
