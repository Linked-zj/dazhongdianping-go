package model

type ShopProduct struct {
	ID          string  `json:"id"`
	ShopId      int     `json:"shopId"`
	ProductName string  `json:"productName"`
	ProductDesc string  `json:"productDesc"`
	OriginPrice float64 `json:"originPrice"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (ShopProduct) TableName() string {
	return "tb_shop_product"
}
