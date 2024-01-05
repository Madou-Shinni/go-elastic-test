package es_doc

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-elastic-client-example/global"
	"log"
	"os"
	"testing"
)

type User struct {
	AccountNumber int `json:"account_number,omitempty"`
	Balance       int `json:"balance,omitempty"`
}

func TestAddDoc(t *testing.T) {
	// 初始化连接
	// 这是必须设置false
	// 设置打印日志 追踪es client执行信息
	logger := log.New(os.Stdout, "ES", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetURL(global.Host), elastic.SetSniff(false), elastic.SetTraceLog(logger))

	if err != nil {
		panic(err)
	}
	// Index a tweet (using JSON serialization)
	tweet1 := User{AccountNumber: 1400, Balance: 48844}
	put1, err := client.Index().
		Index("twitter").
		BodyJson(tweet1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
