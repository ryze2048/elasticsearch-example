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
		Title:     "只要有你",
		Content:   "还珠格格",
		Timestamp: time.Now(),
		Intro:     "没有了你 山河太阳星星都多余",
	}
	if _, err = global.ES.Index().Index(IndexName).BodyJson(article).Do(ctx); err != nil {
		global.ZAPLOG.Error("add data err --> ", zap.Error(err))
		return err
	}

	return nil
}
