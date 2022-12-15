package homework

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/OPengXJ/Homework/interner/repository/mysql/homework"
	"gorm.io/gorm"
)

type CreateHomeworkData struct {
	College   string `form:"college"`
	TeaName   string `form:"teaname"`
	ClassName string `form:"classname"`
	Content   string `form:"content"`
	StartTime string `form:"starttime"`
	Deadline  string `form:"deadline"`
	WorkName  string `form:"workname"`
	Session   string `form:"session"`
}

func (s *Service) Create(data *CreateHomeworkData) error {
	model := homework.NewModel()
	var parseErr error
	model.College = data.College
	model.TeaName = data.TeaName
	model.ClassName = data.ClassName
	model.Content = data.Content
	model.WorkName = data.WorkName
	//这里假设从前端传入的格式为 "2022-12-13T10:11"
	timeFormat := "2006-01-02T15:04"
	startTime, err := time.ParseInLocation(timeFormat, data.StartTime, time.Local)
	if err != nil {
		return err
	}
	deadLine, err := time.ParseInLocation(timeFormat, data.Deadline, time.Local)
	if err != nil {
		return err
	}
	model.StartTime = startTime
	model.Deadline = deadLine
	model.Session, parseErr = strconv.Atoi(data.Session)
	if parseErr != nil {
		return nil
	}
	err = model.Create(s.db.Write)
	if err != nil {
		return err
	}
	//再从mysql读出数据,为了得到id
	cacheData:=homework.NewModel()
	res:=s.db.Write.First(cacheData)
	if res.Error!=nil&&res.Error==gorm.ErrRecordNotFound{
		cacheData=nil
	}
	cache,err:=json.Marshal(cacheData)
	if err!=nil{
		fmt.Println(err.Error())
		return nil
	}
	//作业不经常变动，且发布后经常被学生访问，所以缓存进redis
	//key为学院:年级。 Field为老师:班级  value:作业json数据 过期时间：3/4个作业完成时间
	key := fmt.Sprintf("%s:%d", model.College, model.Session)
	field := fmt.Sprintf("%s:%s", model.TeaName, model.ClassName)
	s.cache.TairRedis.ExHSet(s.ctx, key, field,string(cache))
	//计算缓存的时间
	milesecond:=model.Deadline.Sub(model.StartTime).Milliseconds()
	s.cache.TairRedis.ExHExpire(s.ctx, key, field,int(milesecond))
	return nil
}
