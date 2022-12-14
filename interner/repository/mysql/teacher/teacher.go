package teacher

import (
	"fmt"

	"gorm.io/gorm"
)

func NewModel() *Teacher {
	return new(Teacher)
}



func NewQueryBuilder() *teacherQueryBuilder {
	return new(teacherQueryBuilder)
}

// 将所有的查询条件封装在一个struct中
type teacherQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *teacherQueryBuilder) BuildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	return ret
}

func (qb *teacherQueryBuilder) WhereUsername(value string) *teacherQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "user_name"),
		value,
	})
	return qb
}

func (qb *teacherQueryBuilder) WherePassword(value string) *teacherQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "user_pass"),
		value,
	})
	return qb
}

func (qb *teacherQueryBuilder) WhereTeaCollege(value string) *teacherQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "tea_college"),
		value,
	})
	return qb
}


func (qb *teacherQueryBuilder)Limit(value int)*teacherQueryBuilder{
	qb.limit=value
	return qb
}
func(qb *teacherQueryBuilder)OffSet(value int)*teacherQueryBuilder{
	qb.offset=value
	return qb
}

func (qb *teacherQueryBuilder)Order(value []string)*teacherQueryBuilder{
	qb.order=append(qb.order,value...)
	return qb
}


//经过封装后后的，真正进行相应操作的函数
func (qb *teacherQueryBuilder) First(db *gorm.DB) (*Teacher, error) {
	admin := &Teacher{}
	res := qb.BuildQuery(db).First(admin)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		admin = nil
	}

	return admin, res.Error
}

func (t *Teacher) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


func (qb *teacherQueryBuilder) QueryAll(db *gorm.DB) ([]*Teacher, error) {
	teacher := make([]*Teacher,0)
	res := qb.BuildQuery(db).Find(&teacher)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		teacher = nil
	}
	return teacher, res.Error
}
