package api

import (
	"dazhongdianping-go/global"
	"dazhongdianping-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FlashSaleList(ctx *gin.Context) {
	var flashSaleList []model.FlashSale
	global.DB.Find(&flashSaleList)
	ctx.JSON(http.StatusOK, gin.H{
		"code": global.HTTP_OK,
		"data": flashSaleList,
	})

}
