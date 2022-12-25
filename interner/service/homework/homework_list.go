package homework

import (
	eshomework "github.com/OPengXJ/Homework/interner/repository/elasticsearch/homework"
	"github.com/OPengXJ/Homework/interner/repository/mysql/homework"
)

type SearchHomeworkData struct {
	Page      int      `form:"page"`     //第几页
	PageSize  int      `form:"pagesize"` //每页显示条数
	Session   int      `form:"session"`
	College   string   `form:"college"`
	TeaName   string   `form:"teaname"`
	ClassName string   `form:"classname"`
	Order     []string `form:"order"`
	TeaId     int      `form:"teaid"`
	StuId     int      `form:"stuid"`
}

type ESSearchHomeworkData struct {
	Page      int      `form:"page"`     //第几页
	PageSize  int      `form:"pagesize"` //每页显示条数
	Session   []int      `form:"session"`
	College   []string   `form:"college"`
	TeaName   []string   `form:"teaname"`
	ClassName []string   `form:"classname"`
	Order     []string `form:"order"`
	TeaId     []int      `form:"teaid"`
	StuId     int      `form:"stuid"`
}

func (s *Service) HomeworkList(data *SearchHomeworkData) ([]*homework.Homework, error) {
	qb := homework.NewQueryBuilder()
	if data.Session != 0 {
		qb.WhereSession(data.Session)
	}
	if data.College != "" {
		qb.WhereCollege(data.College)
	}
	if data.ClassName != "" {
		qb.WhereClassName(data.ClassName)
	}
	if data.TeaName != "" {
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
	HomeworkList, err := qb.QueryAll(s.db.Read, s.ctx)
	if err != nil {
		return nil, err
	}
	return HomeworkList, nil
}

// 得到的是已经json化的
// 直接用elasticsearch来实现
func (s *Service) HomeworkListByES(data *ESSearchHomeworkData) (string, error) {
	sb := eshomework.EsNewSearchBuilder(s.es)
	if len(data.Session) != 0 {
		sb.WhereSession(data.Session)
	}
	if len(data.College) != 0 {
		sb.WhereCollege(data.College)
	}
	if len(data.ClassName) != 0 {
		sb.WhereClassName(data.ClassName)
	}
	if len(data.TeaId) != 0 {
		sb.WhereTeaId(data.TeaId)
	}
	if len(data.TeaName) != 0 {
		sb.WhereTeaName(data.TeaName)
	}
	page := data.Page
	if page > 0 {
		page-=1
	}
	pageSize := data.PageSize
	if pageSize == 0 {
		pageSize = 5
	}
	sb.OffSet(page)
	sb.Limit(pageSize)
	if len(data.Order)!=0{
			sb.Order(data.Order)
	}
	sb.BuildSearch()
	return sb.DoSearchHomework(s.ctx)
}
