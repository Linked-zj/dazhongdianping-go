package response

import "dazhongdianping-go/model"

type ShopDetailResponse struct {
	Shop        model.Shop          `json:"shop"`
	Comment     []model.Comment     `json:"comments"`
	ShopProduct []model.ShopProduct `json:"shopProductsList"`
}
