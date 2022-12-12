package token

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/OPengXJ/GoPro/configs"
	"github.com/OPengXJ/Homework/interner/router/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(c *gin.Context, LoginReponse []byte,userType string)  string {
	var rep struct{
		UserId uint `json:"userid"`
		UserName string `json:"username"`
	}
	err:=json.Unmarshal(LoginReponse,&rep)
	if err!=nil{
		c.AbortWithError(http.StatusInternalServerError,err)
	}
	jwtConfig := []byte(configs.Get().JwtPass)
	claims := middlewares.UserClaims{
		UserId:   rep.UserId,
		UserName: rep.UserName,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtConfig)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,err)
	}
	return tokenString
}
