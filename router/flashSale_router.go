package router

import (
	"dazhongdianping-go/api"
	"github.com/gin-gonic/gin"
)

func InitFlashSaleRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("flash")
	{
		ProductRouter.POST("data", api.FlashSaleList)
	}
}
