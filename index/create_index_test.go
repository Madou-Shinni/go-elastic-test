package index

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go-elastic-client-example/global"
	"log"
	"os"
	"testing"
)

const mapping = `
{
	"mappings":{
		"doc":{
			"properties":{
				"name":{
					"type":"text",
					"analyzer": "ik_max_word"
				},
				"id":{
					"type":"integer"
				}
			}
		}
	}
}`

func TestCreateIndex(t *testing.T) {
	// 初始化连接
	// 这是必须设置false
	// 设置打印日志 追踪es client执行信息
	logger := log.New(os.Stdout, "ES", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetURL(global.Host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}
	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("twitter").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("twitter").Body(mapping).IncludeTypeName(true).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}
