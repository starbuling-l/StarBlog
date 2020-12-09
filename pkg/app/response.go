package app

import (
	"github.com/gin-gonic/gin"
	"github.com/starbuling-l/StarBlog/pkg/e"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": make(map[string]string),
	})
	return
}
