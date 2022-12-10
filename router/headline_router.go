package router

import (
	"dazhongdianping-go/api"
	"github.com/gin-gonic/gin"
)

func InitHeadLineRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("swipe")
	{
		ProductRouter.POST("data", api.HeadLineList)
	}
}
