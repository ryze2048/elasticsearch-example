package initialize

import (
	"github.com/ryze2048/elasticsearch-example/global"
	"go.uber.org/zap"
)

func InitLog() {
	global.ZAPLOG = zap.NewExample()
}
