# Homework

>作业提交系统

## 技术栈：

- Gin
- Gorm(Mysql)
- Redis、TairHash moudle
- ElasticSearch
- Zap(log)
- JWT
- Viper

## 简单介绍：

本项目是模拟一个大学的作业布置和完成作业的系统。

1. 使用Viper做一个项目的配置管理。
2. 使用了Gin来做Http框架，负责路由模块和设置JWT的验证中间件；
3. 使用Gorm做Mysql的数据表关系映射和对数据的增删查改的操作；
4. 使用Redis做缓存，目前主要缓存的学生未完成作业的集合；TairHash是一个Redis的module，添加进redis，使redis具备对hash表中的字段进行过期的能力，可以降低空间的消耗，主要是用来存老师刚发布的作业，老师刚发布的作业会被经常访问，且几乎不会修改，所以存于缓存中；
5. 使用ElasticSearch搜索引擎做一个分页筛选排序查询的操作，将作业的标志性信息存于其中，整个分页筛选排序查询操作的逻辑即先选好筛选信息和排序信息，从ElasticSearch中查询，获得排序好的作业列表返回给前端，若想再看更详细的信息，则是根据相应作业的id，先在redis中查询是否有缓存，没有再从mysql中读取信息。
