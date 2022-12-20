package response

type UserLogin struct {
	Token  string `json:"token"`
	UserId int    `json:"userId"`
}
