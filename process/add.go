package process

import (
	"context"
	"github.com/ryze2048/elasticsearch-example/global"
	"go.uber.org/zap"
	"time"
)

type Add struct{}

func (a *Add) AddData(ctx context.Context) (err error) {
	article := Article{
		Title:     "第一天",
		Content:   "无名的人",
		Timestamp: time.Now(),
		Intro:     "毛不易演唱",
	}
	if _, err = global.ES.Index().Index(IndexName).BodyJson(article).Do(ctx); err != nil {
		global.ZAPLOG.Error("add data err --> ", zap.Error(err))
		return err
	}

	return nil
}
