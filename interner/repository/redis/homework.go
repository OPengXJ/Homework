package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)
type HomeworkCreateData struct{
	College string
	ClassName string
	TeaId int
	WorkId	int
	Value string
	ExitTime int
}

type CacheHomworkSearchData struct{
	College string
	ClassName string
	Idlist []string
}

//key为学院:班级。 Field为作业ID  value:作业json数据 过期时间：3/4个作业完成时间
func (c *Cache)SetCache(ctx context.Context,data *HomeworkCreateData)(error){
	key := fmt.Sprintf("%s:%s", data.College, data.ClassName)
	field := fmt.Sprintf("%d",data.WorkId)
	err:=c.TairRedis.ExHSet(ctx, key, field, data.Value).Err()
	if err!=nil{
		return err
	}
	err=c.TairRedis.ExHExpire(ctx, key, field, data.ExitTime).Err()
	if err!=nil{
		//设置过期时间失败，那就要一起先把缓存数据先删除了
		c.TairRedis.ExHDel(ctx,key,field)
		return err
	}
	//相应班级的全体学生未完成作业集合id加上本次作业id
	//在学生模块，可以用来查完成的作业的列表和没完成的列表
	//1.获取本班级所有学生的id
	stulist,err:=c.Redis.SMembers(ctx,data.ClassName+":stulist").Result()
	if err!=nil{
		return err
	}
	//2.给学生未完成作业集合加上作业id
	for _,v:=range stulist{
		class:=fmt.Sprintf("%s:%sundo",data.ClassName,v)
		err=c.Redis.SAdd(ctx,class,data.WorkId).Err()
		if err!=nil{
			return err
		}
	}
	return nil
}


func(c *Cache)TeaWorkId(ctx context.Context,teaid int)([]string,error){
	idList,err:=c.Redis.SMembers(ctx,fmt.Sprintf("TCRHK%d",teaid)).Result()
	if err!=nil{
		return nil,err
	}
	return idList,nil
}

func(c *Cache)StuWorkId(ctx context.Context,classname string)([]string,error){
	idList,err:=c.Redis.SMembers(ctx,fmt.Sprintf("CLSHk%s",classname)).Result()
	if err!=nil{
		return nil,err
	}
	return idList,nil
}

func(c *Cache)WorkInCache(ctx context.Context,data *CacheHomworkSearchData)([]string,[]string,error){
	Result:=make([]string,0)
	idNotExit:=make([]string,0)
	for _,id:=range data.Idlist{
		res,err:=c.TairRedis.ExHGet(ctx,fmt.Sprintf("%s:%s",data.College,data.ClassName),id).Result()
		if err==redis.Nil{
			idNotExit=append(idNotExit, id)
		}else if err!=nil{
			return nil,nil,err
		}
		Result=append(Result, res)
	}
	return Result,idNotExit,nil
}