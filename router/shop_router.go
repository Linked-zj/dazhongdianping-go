package router

import (
	"dazhongdianping-go/api"
	"github.com/gin-gonic/gin"
)

func InitShopRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("shop")
	{
		ProductRouter.POST("list", api.ShopList)
		ProductRouter.POST("detail", api.ShopDetail)

	}
}
