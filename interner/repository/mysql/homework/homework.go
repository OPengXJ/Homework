package homework

import (
	"fmt"

	"gorm.io/gorm"
)

func NewModel() *Homework {
	return new(Homework)
}

func (t *Homework) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewQueryBuilder() *homeworkQueryBuilder {
	return new(homeworkQueryBuilder)
}

// 将所有的查询条件封装在一个struct中
type homeworkQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *homeworkQueryBuilder) BuildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	return ret
}

func (qb *homeworkQueryBuilder) First(db *gorm.DB) (*Homework, error) {
	homework := &Homework{}
	res := qb.BuildQuery(db).First(homework)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		homework = nil
	}

	return homework, res.Error
}

func (qb *homeworkQueryBuilder) WhereClassname(value string) *homeworkQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "classname"),
		value,
	})
	return qb
}
