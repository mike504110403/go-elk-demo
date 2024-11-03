package elk

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

var Es *elasticsearch.Client
var cfg = elasticsearch.Config{
	Addresses: []string{
		"http://localhost:9200",
	},
}

// 初始化
func New(newCfg *elasticsearch.Config) {
	if newCfg != nil {
		cfg = *newCfg
	}
	if es, err := elasticsearch.NewClient(cfg); err != nil {
		panic(err)
	} else {
		Es = es
	}
}

// 新增索引
func IndexRequest(indexreq esapi.IndexRequest) {
	res, err := indexreq.Do(context.Background(), Es)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Println(res)
}

// 查詢索引
func SearchRequest(searchReq esapi.SearchRequest) {
	res, err := searchReq.Do(context.Background(), Es)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Println(res)
}

// 刪除索引
func DeleteRequest(deleteReq esapi.IndicesDeleteRequest) {
	res, err := deleteReq.Do(context.Background(), Es)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Println(res)
}
