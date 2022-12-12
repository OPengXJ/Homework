package student

import (
	"fmt"

	"gorm.io/gorm"
)

func NewModel() *Student {
	return new(Student)
}

func (t *Student) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewQueryBuilder() *studentQueryBuilder {
	return new(studentQueryBuilder)
}

// 将所有的查询条件封装在一个struct中
type studentQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *studentQueryBuilder) BuildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	return ret
}

func (qb *studentQueryBuilder) First(db *gorm.DB) (*Student, error) {
	student := &Student{}
	res := qb.BuildQuery(db).First(student)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		student = nil
	}

	return student, res.Error
}

func (qb *studentQueryBuilder) WhereUsername(value string) *studentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "user_name"),
		value,
	})
	return qb
}

func (qb *studentQueryBuilder) WherePassword(value string) *studentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "user_pass"),
		value,
	})
	return qb
}

func (qb *studentQueryBuilder) WhereStuClass(value string) *studentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "stu_class"),
		value,
	})
	return qb
}


