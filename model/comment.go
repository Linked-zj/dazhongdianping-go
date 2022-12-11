package model

type Comment struct {
	ID                 int    `json:"id"`
	UserId             int    `json:"userId"`
	UserName           string `json:"userName"`
	UserAvatarImg      string `json:"userAvatarImg"`
	ShopCommentedId    int    `json:"shopCommentedId"`
	ProductCommentedId int    `json:"productCommentedId"`
	IfComment          int    `json:"ifComment"`
	CommentInfo        string `json:"commentInfo"`
	CommentImg1        string `json:"commentImg1"`
	CommentImg2        string `json:"commentImg2"`
	CommentImg3        string `json:"commentImg3"`
	CommentImg4        string `json:"commentImg4"`
	CommentImg5        string `json:"commentImg5"`
	CommentImg6        string `json:"commentImg6"`
	CreateTime         string `json:"createTime"`
}

func (Comment) TableName() string {
	return "tb_comment"
}
