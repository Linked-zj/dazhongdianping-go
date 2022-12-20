package model

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Telephone  string    `json:"telephone"`
	Password   string    `json:"password"`
	Gender     int       `json:"gender"`
	Role       int       `json:"role"`
	CreateTime time.Time `json:"createTime"`
	Birthday   string    `json:"birthday"`
}

func (User) TableName() string {
	return "tb_user"
}
