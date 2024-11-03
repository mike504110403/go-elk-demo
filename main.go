package main

import (
	"go-elk-demo/elk"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

func Init() {
	elk.New(&elasticsearch.Config{})
}

func main() {
	Init()
	// 新增索引
	elk.IndexRequest(esapi.IndexRequest{
		Index: "test_index",
		Body:  strings.NewReader(`{"title":"go es index test", "Content": "elasticsearch client"}`),
	})
	// 查詢索引
	elk.SearchRequest(esapi.SearchRequest{
		Index: []string{"test_index"},
		Body:  strings.NewReader(`{"from":0,"query":{"bool":{"must":[{"match":{"title":{"operator":"and","query":"go es index test"}}}]}},"size":20}`),
	})
	// 刪除索引
	elk.DeleteRequest(esapi.IndicesDeleteRequest{
		Index: []string{"test_index"},
	})
}
