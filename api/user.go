package api

import (
	"crypto/md5"
	"dazhongdianping-go/global"
	"dazhongdianping-go/middlewares"
	"dazhongdianping-go/model"
	"dazhongdianping-go/model/request"
	"dazhongdianping-go/model/response"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func getMd5String(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func UserLogin(ctx *gin.Context) {
	var user model.User
	var userLogin request.UserLogin
	if err := ctx.BindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	enpPwd := getMd5String(userLogin.Password)
	fmt.Println(enpPwd)

	err := global.DB.Where("telephone = ? and password = ? ", userLogin.LoginName, enpPwd).Find(&user).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": "用户不存在",
		})
		return
	}
	uuid := uuid.New().String()
	tokenStr := strings.Replace(uuid, "-", "", -1)
	j := middlewares.NewJWT()
	claims := model.CustomClaims{
		UUID: tokenStr,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24, // 过期时间
			Issuer:    "Linked",
		},
	}
	token, err := j.CreateToken(claims)
	var userLoginRep response.UserLogin
	userLoginRep.Token = token
	userLoginRep.UserId = user.ID
	ctx.JSON(http.StatusOK, gin.H{
		"code": global.HTTP_OK,
		"data": userLoginRep,
	})
}
