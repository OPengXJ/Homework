package homework

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func NewModel() *Homework {
	return new(Homework)
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
	ret.Limit(qb.limit)
	ret.Offset(qb.offset)
	for _,order:=range qb.order{
		ret=ret.Order(order)
	}
	return ret
}

func (qb *homeworkQueryBuilder) WhereCollege(value string) *homeworkQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "college"),
		value,
	})
	return qb
}

func (qb *homeworkQueryBuilder) WhereSession(value int) *homeworkQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "session"),
		value,
	})
	return qb
}

func (qb *homeworkQueryBuilder) WhereClassName(value string) *homeworkQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "class_name"),
		value,
	})
	return qb
}

func (qb *homeworkQueryBuilder) WhereTeaName(value string) *homeworkQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "tea_name"),
		value,
	})
	return qb
}

func (qb *homeworkQueryBuilder)Limit(value int)*homeworkQueryBuilder{
	qb.limit=value
	return qb
}
func(qb *homeworkQueryBuilder)OffSet(value int)*homeworkQueryBuilder{
	qb.offset=value
	return qb
}

func (qb *homeworkQueryBuilder)Order(value []string)*homeworkQueryBuilder{
	qb.order=append(qb.order,value...)
	return qb
}
func (qb *homeworkQueryBuilder) WhereIdIn(value []string) *homeworkQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}



//封装后，最终调用的方法
func (t *Homework) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (qb *homeworkQueryBuilder) First(db *gorm.DB,ctx context.Context) (*Homework, error) {
	homework := &Homework{}
	res := qb.BuildQuery(db).WithContext(ctx).First(homework)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		homework = nil
	}

	return homework, res.Error
}

func (qb *homeworkQueryBuilder) QueryAll(db *gorm.DB,ctx context.Context) ([]*Homework, error) {
	homework := make([]*Homework,0)
	res := qb.BuildQuery(db).WithContext(ctx).Find(&homework)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		homework = nil
	}
	return homework, res.Error
}
