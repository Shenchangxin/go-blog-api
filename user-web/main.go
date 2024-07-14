package main

import (
	"fmt"
	"go-blog-api/user-web/initialize"
	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
	Routers := initialize.Routers()

	zap.S().Infof("启动服务器：%d", 8021)

	if err := Routers.Run(fmt.Sprintf(":%d", 8021)); err != nil {
		zap.S().Panicf("启动服务器：%d 失败", 8021)
	}

}
