package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger
var suggerLogger *zap.SugaredLogger

func InitLogger() {
	//logger, _ = zap.NewProduction()
	logger, _ = zap.NewDevelopment()
	//返回json格式

	suggerLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		suggerLogger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))
	} else {
		suggerLogger.Info("Success..",
			zap.String("stateCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
