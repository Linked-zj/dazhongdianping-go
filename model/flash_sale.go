package model

type FlashSale struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	OriginPrice float64 `json:"originPrice"`
	ImgUrl      string  `json:"imgUrl"`
	Spec        string  `json:"spec"`
	Count       int     `json:"count"`
	Sales       int     `json:"sales"`
	Status      int     `json:"status"`
}

func (FlashSale) TableName() string {
	return "tb_flash_sale"
}
