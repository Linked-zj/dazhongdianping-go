package main

import (
	"dazhongdianping-go/global"
	"dazhongdianping-go/initialize"
	"dazhongdianping-go/router"
	"fmt"
)

func main() {
	// 2. 初始化配置
	initialize.InitConfig()
	Router := router.Routers()

	initialize.InitGorm()

	Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
}
