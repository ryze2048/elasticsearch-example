package initialize

import (
	"github.com/olivere/elastic/v7"
	"github.com/ryze2048/elasticsearch-example/global"
	"go.uber.org/zap"
)

func InitElasticsearch() {
	var err error
	var client *elastic.Client
	if client, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false)); err != nil {
		global.ZAPLOG.Error("init es client err --> ", zap.Error(err))
		return
	}

	global.ES = client
	global.ZAPLOG.Info("init es client successful")
	return
}
