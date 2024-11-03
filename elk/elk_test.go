package elk

import (
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

func Init() {
	New(&elasticsearch.Config{})
}

func TestIndexRequest(t *testing.T) {
	Init()
	// 新增索引
	IndexRequest(esapi.IndexRequest{
		Index: "test_index",
		Body:  strings.NewReader(`{"title":"go es index test", "Content": "elasticsearch client"}`),
	})
}

func TestSearchRequest(t *testing.T) {
	Init()
	// 查詢索引
	SearchRequest(esapi.SearchRequest{
		Index: []string{"test_index"},
		Body:  strings.NewReader(`{"from":0,"query":{"bool":{"must":[{"match":{"title":{"operator":"and","query":"go es index test"}}}]}},"size":20}`),
	})
}

func TestDeleteRequest(t *testing.T) {
	Init()
	// 刪除索引
	DeleteRequest(esapi.IndicesDeleteRequest{
		Index: []string{"test_index"},
	})
}
