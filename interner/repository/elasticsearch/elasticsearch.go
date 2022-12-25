package elasticsearch

import (
	"time"

	"github.com/OPengXJ/GoPro/pkg/log"
	"github.com/OPengXJ/Homework/configs"
	"github.com/olivere/elastic/v7"
)

var esClient = new(elastic.Client)

func init() {
	config := configs.Get().ElasticSearch
	client, err := elastic.NewClient(
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL(config.Addr),
		elastic.SetSniff(false),
		// 基于http base auth验证机制的账号和密码
		elastic.SetBasicAuth(config.User, config.Pass),
		// 启用gzip压缩
		elastic.SetGzip(config.Gzip),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(config.CheckInterval*time.Second),
		// 设置错误日志输出
		elastic.SetErrorLog(log.NewEsError()),
		// 设置info日志输出
		elastic.SetInfoLog(log.NewEsInfo()),
	)

	if err != nil {
		// Handle error
		panic(err)
	}
	esClient = client

}

func GetEsClient() *elastic.Client {
	return esClient
}
