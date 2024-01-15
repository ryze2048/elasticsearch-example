package main

import (
	"context"
	"fmt"
	"github.com/ryze2048/elasticsearch-example/global"
	"github.com/ryze2048/elasticsearch-example/initialize"
	"github.com/ryze2048/elasticsearch-example/process"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.InitLog()
	initialize.InitElasticsearch()
	ctx, cancel := context.WithCancel(context.Background())

	var p process.Process
	p.Index.Create(ctx)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case sig := <-sigChan:
			global.ZAPLOG.Info("To exit press CTRL+C")
			// 处理中断信号
			global.ZAPLOG.Info(fmt.Sprintf("Received signal: %s", sig))
			cancel()
			os.Exit(0)
		}
	}
}
