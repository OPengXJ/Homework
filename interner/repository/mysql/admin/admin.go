package admin

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func NewModel() *Admin {
	return new(Admin)
}

func NewQueryBuilder() *adminQueryBuilder {
	return new(adminQueryBuilder)
}

// 将所有的查询条件封装在一个struct中
type adminQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *adminQueryBuilder) BuildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	return ret
}

func (qb *adminQueryBuilder) WhereUsername(value string) *adminQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "user_name"),
		value,
	})
	return qb
}

func (qb *adminQueryBuilder) WherePassword(value string) *adminQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v = ?", "user_pass"),
		value,
	})
	return qb
}


//封装后真正调用的函数

func (t *Admin) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (qb *adminQueryBuilder) First(db *gorm.DB,ctx context.Context) (*Admin, error) {
	admin := &Admin{}
	res := qb.BuildQuery(db).WithContext(ctx).First(admin)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		admin = nil
	}

	return admin, res.Error
}