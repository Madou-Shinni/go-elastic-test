package query

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-elastic-client-example/global"
	"testing"
)

func TestMatchQuery(t *testing.T) {
	// 初始化连接
	// 这是必须设置false
	client, err := elastic.NewClient(elastic.SetURL(global.Host), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	query := elastic.NewMatchQuery("address", "street")

	res, err := client.Search().Index("users").Query(query).Do(context.Background())
	if err != nil {
		panic(err)
	}

	total := res.Hits.TotalHits.Value
	fmt.Printf("搜索结果数量：%v\n", total)

	for _, hit := range res.Hits.Hits {
		var user User
		err := json.Unmarshal(hit.Source, &user)
		if err != nil {
			// Deserialization failed
		}

		fmt.Println(user)
	}
}
