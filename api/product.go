package api

import (
	"dazhongdianping-go/global"
	"dazhongdianping-go/model"
	"dazhongdianping-go/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProductList(ctx *gin.Context) {
	var p []model.Product
	var productList request.ProductListById
	if err := ctx.BindJSON(&productList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	global.DB.Where("category_id = ?", productList.CategoryId).Find(&p)
	ctx.JSON(http.StatusOK, gin.H{
		"code": global.HTTP_OK,
		"data": p,
	})
}

func ProductFlashList(ctx *gin.Context) {
	var p []model.Product
	global.DB.Where("is_flash = ?", 1).Find(&p)
	ctx.JSON(http.StatusOK, gin.H{
		"code": global.HTTP_OK,
		"data": p,
	})
}
