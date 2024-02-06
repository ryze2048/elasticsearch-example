package process

import (
	"context"
	"github.com/ryze2048/elasticsearch-example/global"
	"go.uber.org/zap"
)

type Index struct{}

// Create 创建索引
func (i *Index) Create(ctx context.Context) {
	var err error
	if _, err = global.ES.CreateIndex(IndexName).BodyString(indexMapping).Do(ctx); err != nil {
		global.ZAPLOG.Error("create es index err --> ", zap.Error(err))
	}

	global.ZAPLOG.Info("create es index successful")
}
