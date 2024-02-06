package process

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/ryze2048/elasticsearch-example/global"
	"go.uber.org/zap"
	"time"
)

type Update struct{}

func (u *Update) Update(ctx context.Context) (err error) {
	updateByQuery := elastic.NewUpdateByQueryService(global.ES)

	updateByQuery.Index(IndexName).Query(elastic.NewMatchQuery("title", "第一天")).
		Script(elastic.NewScriptInline("ctx._source.timestamp = params.new_timestamp").Lang("painless").Param("new_timestamp", time.Now()))

	var response *elastic.BulkIndexByScrollResponse
	if response, err = updateByQuery.Do(ctx); err != nil {
		global.ZAPLOG.Error("update index err --> ", zap.Error(err))
		return err
	}

	global.ZAPLOG.Info(fmt.Sprintf("update %d documents", response.Updated))
	return nil
}
