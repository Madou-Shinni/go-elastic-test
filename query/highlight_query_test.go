package query

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-elastic-client-example/global"
	"log"
	"os"
	"testing"
)

func TestHighlightQuery(t *testing.T) {
	// 初始化连接
	// 这是必须设置false
	// 设置打印日志 追踪es client执行信息
	logger := log.New(os.Stdout, "ES", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetURL(global.Host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Specify highlighter
	hl := elastic.NewHighlight()
	hl = hl.Fields(elastic.NewHighlighterField("address"))
	hl = hl.PreTags("<em>").PostTags("</em>")

	// Match all should return all documents
	query := elastic.NewPrefixQuery("address", "street")

	// 执行搜索请求
	searchResult, err := client.Search().
		Index("users").
		Highlight(hl).
		Query(query).
		Pretty(true).
		Do(context.TODO())
	if err != nil {
		log.Fatalf("Error executing search request: %s", err)
	}

	var (
		List     []interface{}
		userList []User
	)

	// 处理搜索结果
	for _, hit := range searchResult.Hits.Hits {
		// 打印文档ID
		fmt.Printf("Document ID: %s\n", hit.Id)

		// 打印高亮结果
		if len(hit.Highlight) > 0 {
			hightStr := hit.Highlight["address"][0]
			fmt.Printf("Highlight: %s\n", hightStr)

			// 替换高亮结果
			bytes, err := hit.Source.MarshalJSON()
			if err != nil {
				continue
			}

			// 反序列化结果至map
			m := make(map[string]interface{})
			err = json.Unmarshal(bytes, &m)
			if err != nil {
				log.Fatalf("json.Unmarshal error: %s", err.Error())
			}

			// 替换原是文本为高亮结果
			m["address"] = hightStr
			List = append(List, m)

			// 将map序列化
			str, err := json.Marshal(m)
			if err != nil {
				log.Fatalf("json.Marshal error: %s", err.Error())
			}

			// 反序列化结果到user
			var user User
			err = json.Unmarshal(str, &user)
			if err != nil {
				log.Fatalf("json.Unmarshal error: %s", err.Error())
			}

			userList = append(userList, user)
		}
	}

	for _, item := range List {
		log.Println(item)
	}

	log.Println(userList)
}
