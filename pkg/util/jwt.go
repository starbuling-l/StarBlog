package util

import "github.com/starbuling-l/StarBlog/pkg/setting"

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
