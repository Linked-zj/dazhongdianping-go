package api

import (
	"dazhongdianping-go/global"
	"dazhongdianping-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentSelectList(ctx *gin.Context) {
	var comment []model.Comment
	global.DB.Where("if_comment", 1).Find(&comment)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": comment,
	})
}
