package class

import (
	"fmt"

	"gorm.io/gorm"
)

func NewModel() *Class {
	return new(Class)
}

func (t *Class) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewQueryBuilder() *classQueryBuilder {
	return new(classQueryBuilder)
}

// 将所有的查询条件封装在一个struct中
type classQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *classQueryBuilder) BuildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	ret.Limit(qb.limit)
	ret.Offset(qb.offset)
	for _,order:=range qb.order{
		ret=ret.Order(order)
	}
	return ret
}

func (qb *classQueryBuilder) First(db *gorm.DB) (*Class, error) {
	class := &Class{}
	res := qb.BuildQuery(db).First(class)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		class = nil
	}

	return class, res.Error
}

func (qb *classQueryBuilder) WhereCollege(value string) *classQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "college"),
		value,
	})
	return qb
}

func (qb *classQueryBuilder) WhereSession(value int) *classQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "session"),
		value,
	})
	return qb
}

func (qb *classQueryBuilder)Limit(value int)*classQueryBuilder{
	qb.limit=value
	return qb
}
func(qb *classQueryBuilder)OffSet(value int)*classQueryBuilder{
	qb.offset=value
	return qb
}

func (qb *classQueryBuilder)Order(value []string)*classQueryBuilder{
	qb.order=append(qb.order,value...)
	return qb
}

func (qb *classQueryBuilder) QueryAll(db *gorm.DB) ([]*Class, error) {
	class := make([]*Class,0)
	res := qb.BuildQuery(db).Find(&class)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		class = nil
	}
	return class, res.Error
}
