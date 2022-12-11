package router

import (
	"dazhongdianping-go/api"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("comment")
	{
		ProductRouter.POST("select", api.CommentSelectList)

	}
}
