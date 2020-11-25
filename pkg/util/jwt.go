package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, passward string) (string, error) {
	expiresTime := time.Now().Add(3 * time.Hour)
	claims := Claims{
		username,
		passward,
		jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
			Issuer:    "gin",
		}}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
