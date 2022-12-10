package api

import (
	"dazhongdianping-go/global"
	"dazhongdianping-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HeadLineList(ctx *gin.Context) {
	var headImgList []model.HeadLine
	global.DB.Find(&headImgList)
	ctx.JSON(http.StatusOK, gin.H{
		"code": global.HTTP_OK,
		"data": headImgList,
	})

}
