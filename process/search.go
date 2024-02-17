package process

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/ryze2048/elasticsearch-example/global"
	"go.uber.org/zap"
	"log"
)

type Search struct{}

func (s *Search) Search(ctx context.Context) {
	searchResult, err := global.ES.Search().
		Index(IndexName).
		Query(elastic.NewMatchQuery("intro", "鸟")).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error executing search: %s", err)
	}

	// 处理搜索结果
	if searchResult.Hits.TotalHits.Value > 0 {
		global.ZAPLOG.Info(fmt.Sprintf("Found %d hits:", searchResult.Hits.TotalHits.Value))
		for _, hit := range searchResult.Hits.Hits {
			var doc Article
			if err = json.Unmarshal(hit.Source, &doc); err != nil {
				global.ZAPLOG.Error("json unmarshal err --> ", zap.Error(err))
				return
			}
			fmt.Println("doc --> ", doc)
		}

	} else {
		log.Println("No hits found")
	}
}

func (s *Search) SearchAnd(ctx context.Context) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewMatchQuery("title", "你"))
	boolQuery.Must(elastic.NewMatchQuery("intro", "星星"))

	searchResult, err := global.ES.Search().Index(IndexName).Query(boolQuery).Do(ctx)
	if err != nil {
		global.ZAPLOG.Error("search err --> ", zap.Error(err))
		return
	}

	if searchResult.Hits.TotalHits.Value > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var doc Article
			if err = json.Unmarshal(hit.Source, &doc); err != nil {
				global.ZAPLOG.Error("json unmarshal err --> ", zap.Error(err))
				return
			}
			fmt.Println("doc --> ", doc)
		}
	}
}

func (s *Search) SearchOr(ctx context.Context) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Should(elastic.NewMatchQuery("title", "你"))
	boolQuery.Should(elastic.NewMatchQuery("intro", "星星"))

	searchResult, err := global.ES.Search().Index(IndexName).Query(boolQuery).Do(ctx)
	if err != nil {
		global.ZAPLOG.Error("search err --> ", zap.Error(err))
		return
	}

	if searchResult.Hits.TotalHits.Value > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var doc Article
			if err = json.Unmarshal(hit.Source, &doc); err != nil {
				global.ZAPLOG.Error("json unmarshal err --> ", zap.Error(err))
				return
			}
			fmt.Println("doc --> ", doc)
		}
	}
}
