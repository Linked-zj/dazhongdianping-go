package router

import (
	"dazhongdianping-go/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {

	Router := gin.Default()

	Router.Use(middlewares.Cors())
	ApiRouter := Router.Group("o2o")
	ApiRouter.StaticFS("/image/", http.Dir("./resource/images/dazhongdianping"))
	InitProductRouter(ApiRouter)
	InitHeadLineRouter(ApiRouter)
	InitFlashSaleRouter(ApiRouter)
	return Router

}
