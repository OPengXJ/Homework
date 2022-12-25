package elasticsearch

import "time"

type EsHomeworkSearch struct{
	Id	int `json:"id"`	//设置文档id
	StartTime time.Time	`json:"starttime"`//放一个时间戳，后面有根据时间排序的需求
	DeadLine	time.Time	`json:"deadline"`//放一个时间戳，后面有根据时间排序的需求
	ClassName	string	`json:"classname"`//后面会有根据老师筛选数据的需求，和根据班级名称排序的需求
	TeaName	string	`json:"teaname"`//后面会有根据老师筛选数据的需求，和根据老师名称排序的需求
	College	string	`json:"college"`//同上
	Session	int	`json:"session"`//一样
	TeaId	int	`json:"teaid"`//考虑是在，老师页面时，可以根据老师的id来筛选
	WorkId	int	`json:"workid"`	//最重要的，到时就靠这个id去拿数据，和上面的id其实时一样的
}


