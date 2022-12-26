package homework

import (
	"encoding/json"
	"fmt"

	"github.com/OPengXJ/Homework/interner/repository/mysql/homework"
)

type HomeworkDetailSearchData struct{
	//主要对应的是在redis中的key和value
	Id	string	`form:"id"`
	College string	`form:"college"`
	ClassName	string	`form:"classname"`

}
func(s *Service)HomeworkDeatil(data *HomeworkDetailSearchData)(string,error){
	res:=s.cache.TairRedis.ExHGet(s.ctx,fmt.Sprintf("%s:%s",data.College,data.ClassName),data.Id)
	result,err:=res.Result()
	if err==nil{
		fmt.Println("get from cache")
		return result,nil
	}
	qb:=homework.NewQueryBuilder()
	qb.WhereId(data.Id)
	qb.BuildQuery(s.db.Read)
	dbData,err:=qb.First(s.db.Read,s.ctx)
	if err!=nil{
		return "",err
	}
	jsonData,err:=json.Marshal(dbData)
	if err!=nil{
		return "",err
	}
	return string(jsonData),nil
}