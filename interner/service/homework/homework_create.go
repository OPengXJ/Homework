package homework

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/OPengXJ/Homework/interner/repository/elasticsearch"
	eshomework "github.com/OPengXJ/Homework/interner/repository/elasticsearch/homework"
	"github.com/OPengXJ/Homework/interner/repository/mysql/homework"
	"github.com/OPengXJ/Homework/interner/repository/redis"
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
	TeaId	int	`form:"teaid"`
}

func (s *Service) Create(data *CreateHomeworkData) error {
	model := homework.NewModel()
	var parseErr error
	model.College = data.College
	model.TeaName = data.TeaName
	model.ClassName = data.ClassName
	model.Content = data.Content
	model.WorkName = data.WorkName
	model.TeaId=data.TeaId
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
	cache, err := json.Marshal(model)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//作业不经常变动，且发布后经常被学生访问，所以缓存进redis
	//计算缓存的时间
	milesecond := model.Deadline.Sub(model.StartTime).Milliseconds()*3/4
	setdata:=&redis.HomeworkCreateData{
		College: model.College,
		ClassName: model.ClassName,
		TeaId: model.TeaId,
		Value: string(cache),
		WorkId: int(model.ID),
		ExitTime: int(milesecond),
	}
	err=s.cache.SetCache(s.ctx,setdata)
	if err!=nil{
		fmt.Println("设置缓存失败",err.Error())
		return err
	}
	esHomeWorkSearch:=eshomework.NewModel()
	esHomeWorkSearch.Id=int(model.ID)
	esHomeWorkSearch.StartTime=model.StartTime
	esHomeWorkSearch.DeadLine=model.Deadline
	esHomeWorkSearch.ClassName=model.ClassName
	esHomeWorkSearch.TeaName=model.TeaName
	esHomeWorkSearch.College=model.College
	esHomeWorkSearch.Session=model.Session
	esHomeWorkSearch.TeaId=model.TeaId
	esHomeWorkSearch.WorkId=int(model.ID)
	err=esHomeWorkSearch.Create(elasticsearch.GetEsClient(),s.ctx)
	if err!=nil{
		return err
	}
	return nil
}
