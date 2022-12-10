package request

type ProductListById struct {
	CategoryId int `json:"categoryId" form:"categoryId" binding:"required"`
}
