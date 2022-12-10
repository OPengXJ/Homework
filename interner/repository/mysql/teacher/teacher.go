package teacher

import (
	"fmt"

	"gorm.io/gorm"
)

func NewModel() *Teacher {
	return new(Teacher)
}

func (t *Teacher) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
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

func (qb *teacherQueryBuilder) First(db *gorm.DB) (*Teacher, error) {
	admin := &Teacher{}
	res := qb.BuildQuery(db).First(admin)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		admin = nil
	}

	return admin, res.Error
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
