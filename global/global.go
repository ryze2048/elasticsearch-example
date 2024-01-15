package global

import (
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

var ZAPLOG *zap.Logger

var ES *elastic.Client
