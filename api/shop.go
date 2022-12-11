package api

import (
	"dazhongdianping-go/global"
	"dazhongdianping-go/model"
	"dazhongdianping-go/model/request"
	"dazhongdianping-go/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShopList(ctx *gin.Context) {
	var shops []model.Shop
	var shopList request.ShopList
	if err := ctx.BindJSON(&shopList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	if shopList.CategoryId != 0 {
		global.DB.Where("shop_category_id = ?", shopList.CategoryId).Find(&shops)
	} else {
		global.DB.Find(&shops)

	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": global.HTTP_OK,
		"data": shops,
	})
}

func ShopDetail(ctx *gin.Context) {
	var shopDatailRes response.ShopDetailResponse
	var shopDetail request.ShopDetail
	if err := ctx.BindJSON(&shopDetail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	var shop model.Shop
	var comment []model.Comment
	var shopProduct []model.ShopProduct
	global.DB.First(&shop, shopDetail.ShopId)
	global.DB.Where("shop_commented_id", shopDetail.ShopId).Find(&comment)
	global.DB.Where("shop_id", shopDetail.ShopId).Find(&shopProduct)
	shopDatailRes.Shop = shop
	shopDatailRes.Comment = comment
	shopDatailRes.ShopProduct = shopProduct
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": shopDatailRes,
	})
}
