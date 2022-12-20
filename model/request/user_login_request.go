package request

type UserLogin struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}
