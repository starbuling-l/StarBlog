package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/starbuling-l/StarBlog/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, passward string) (string error) {
	expiresTime = tie
		claims := Claims{
		username,
		passward,
		jwt.StandardClaims{
			ExpiresAt:,
			Issuer: "gin",
		}}
	return ,error()
}
