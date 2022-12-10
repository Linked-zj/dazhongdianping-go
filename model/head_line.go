package model

type HeadLine struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	ImgUrl string `json:"imgUrl"`
}

func (HeadLine) TableName() string {
	return "tb_head_line"
}
