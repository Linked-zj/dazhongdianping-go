package router

import (
	"dazhongdianping-go/api"
	"github.com/gin-gonic/gin"
)

func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("product")
	{
		ProductRouter.POST("list", api.ProductList)
		ProductRouter.POST("flash", api.ProductFlashList)

	}
}
