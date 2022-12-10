package model

type Product struct {
	ID          int     `json:"id"`
	CategoryId  int     `json:"categoryId"`
	Name        string  `json:"name"`
	OriginPrice float64 `json:"originPrice"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
	IsFlash     int     `json:"isFlash"`
	Count       int     `json:"count"`
	Status      int     `json:"status"`
	Sales       int     `json:"sales"`
}

func (Product) TableName() string {
	return "tb_product"
}
