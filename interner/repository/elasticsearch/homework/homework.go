package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"github.com/olivere/elastic/v7"
)

func NewModel() *EsHomeworkSearch {
	return new(EsHomeworkSearch)
}

type HomeworkSearchBuilder struct {
	order []string
	whereNo []struct {
		prefix string
		stringValue  []string
		intValue	[]int
	}
	where []struct {
		prefix string
		stringValue  []string
		intValue	[]int
	}
	limit         int
	offset        int
	SearchService *elastic.SearchService
}

func EsNewSearchBuilder(esclient *elastic.Client) *HomeworkSearchBuilder {
	return &HomeworkSearchBuilder{
		SearchService: esclient.Search(),
	}
}
func(cb *HomeworkSearchBuilder)Limit(value int){
	cb.limit=value
}
func(cb *HomeworkSearchBuilder)OffSet(value int){
	cb.offset=value
}
func(cb *HomeworkSearchBuilder)Order(value []string){
	cb.order=value
}
func (cb *HomeworkSearchBuilder)WhereNoId(value []int) {
	cb.whereNo = append(cb.whereNo, struct {
		prefix string
		stringValue  []string
		intValue	[]int
	}{
		prefix:"_id",
		intValue: value,
	})
}

func (cb *HomeworkSearchBuilder) WhereClassName(value []string) {
	cb.where = append(cb.where, struct {
		prefix string
		stringValue  []string
		intValue	[]int
	}{
		prefix:"classname",
		stringValue:value,
	})
}

func (cb *HomeworkSearchBuilder) WhereTeaName(value []string) {
	cb.where = append(cb.where, struct {
		prefix string
		stringValue  []string
		intValue	[]int
	}{
		prefix:"teaname",
		stringValue:value,
	})
}

func (cb *HomeworkSearchBuilder) WhereCollege(value []string) {
	cb.where = append(cb.where, struct {
		prefix string
		stringValue  []string
		intValue	[]int
	}{
		prefix:"college",
		stringValue: value,
	})
}

func (cb *HomeworkSearchBuilder) WhereSession(value []int) {
	cb.where = append(cb.where, struct {
		prefix string
		stringValue  []string
		intValue	[]int
	}{
		prefix:"session",
		intValue:value,
	})
}

func (cb *HomeworkSearchBuilder) WhereTeaId(value []int) {
	cb.where = append(cb.where, struct {
		prefix string
		stringValue  []string
		intValue	[]int
	}{
		prefix:"teaid",
		intValue:value,
	})
}



func (es *EsHomeworkSearch)Create(esclient *elastic.Client,ctx context.Context)error{
	_,err:=esclient.Index().
	Index("homework").
	Id(strconv.Itoa(es.Id)).
	BodyJson(es).Do(ctx)
	if err!=nil{
		return err
	}
	return nil
}
func (cb *HomeworkSearchBuilder) BuildSearch() {
	boolQuery := elastic.NewBoolQuery()
	for _, termQuery := range cb.where {
		if len(termQuery.intValue)!=0{
			tempBool:=elastic.NewBoolQuery()
			for _,v:=range termQuery.intValue{
				tempBool.Should(elastic.NewTermsQuery(termQuery.prefix,v))
			}
			boolQuery=boolQuery.Must(tempBool)
		}else{
			tempBool:=elastic.NewBoolQuery()
			for _,v:=range termQuery.stringValue{
				tempBool.Should(elastic.NewTermsQuery(termQuery.prefix+".keyword",v))
			}
			boolQuery=boolQuery.Must(tempBool)
			fmt.Println(termQuery.stringValue)
		}
	}
	for _,termQueryNo:=range cb.whereNo{
		//
		if len(termQueryNo.intValue)!=0{
			tempBool:=elastic.NewBoolQuery()
			for _,v:=range termQueryNo.intValue{
				tempBool.Should(elastic.NewTermsQuery(termQueryNo.prefix,v))
			}
			boolQuery=boolQuery.Must(tempBool)
		}else{
			tempBool:=elastic.NewBoolQuery()
			for _,v:=range termQueryNo.stringValue{
				tempBool.Should(elastic.NewTermsQuery(termQueryNo.prefix+".keyword",v))
			}
			boolQuery=boolQuery.Must(tempBool)
			fmt.Println(termQueryNo.stringValue)
		}
		//
	}
	cb.SearchService = cb.SearchService.Index("homework").Query(boolQuery)
	for _, order := range cb.order {
		fmt.Println("order",order)
		cb.SearchService = cb.SearchService.SortBy(elastic.NewFieldSort(order+".keyword").Desc())
	}
	cb.SearchService = cb.SearchService.From(cb.offset).Size(cb.limit)
}

func (cb *HomeworkSearchBuilder) DoSearchHomework(ctx context.Context) (string,error) {

	result, err := cb.SearchService.
	Pretty(true).
	Do(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return "",err
	}
	fmt.Println("total:",result.TotalHits())
	fmt.Println(result.Each(reflect.TypeOf(&EsHomeworkSearch{}))...)
	jsonData, err := json.Marshal(result.Hits.Hits)
	if err != nil {
		return "",err
	}
	return string(jsonData),nil
}
