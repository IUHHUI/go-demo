package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	url := "http://afonddream.com"
	// 结构化日志打印
	logger.Sugar().Infow("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)

	// 非结构化日志打印
	logger.Sugar().Infof("failed to fetch URL: %s", url)

	/*
		{"level":"info","ts":1627909431.181877,"caller":"log-zap/main.go:14","msg":"failed to fetch URL","url":"http://afonddream.com","attempt":3,"backoff":1}
		{"level":"info","ts":1627909431.182239,"caller":"log-zap/main.go:17","msg":"failed to fetch URL: http://afonddream.com"}
	*/
}
