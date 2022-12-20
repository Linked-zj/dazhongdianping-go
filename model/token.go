package model

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UUID string
	//ID          uint
	//NickName    string
	//AuthorityId uint
	jwt.StandardClaims
}
