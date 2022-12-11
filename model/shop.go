package model

type Shop struct {
	ID             int     `json:"id" gorm:"primary_key"`
	ShopCategoryId int     `json:"shopCategoryId"`
	ShopAreaId     string  `json:"shopAreaId"`
	Pritory        int     `json:"pritory"`
	ShopName       string  `json:"shopName"`
	ShopDesc       string  `json:"shopDesc"`
	ShopAvgPrice   int     `json:"shopAvgPrice"`
	ShopPhone      string  `json:"shopPhone"`
	ShopAddress    string  `json:"shopAddress"`
	ShopRate       float64 `json:"shopRate"`
	CreateTime     string  `json:"createTime"`
	LastEditTime   string  `json:"lastEditTime"`
	ShopLongitude  string  `json:"shopLongitude"`
	ShopLatitude   string  `json:"shopLatitude"`
	ShopImage1     string  `json:"shopImage1"`
	ShopImage2     string  `json:"shopImage2"`
	ShopImage3     string  `json:"shopImage3"`
}

func (Shop) TableName() string {
	return "tb_shop"
}
