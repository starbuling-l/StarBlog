package api

import (
	"net/http"

	"github.com/starbuling-l/StarBlog/models"
	"github.com/starbuling-l/StarBlog/pkg/app"
	"github.com/starbuling-l/StarBlog/pkg/e"
	"github.com/starbuling-l/StarBlog/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

//获取token curl 127.0.0.1:9000/api/auth
// @Summary 获取token
// @Produce  json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/auth [get]
func GetAuth(c *gin.Context) {
	g := app.Gin{C: c}
	/*	data := make(map[string]interface{})
		code := e.INVALID_PARAMS*/
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := auth{
		Username: username,
		Password: password,
	}

	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		g.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	exist, err := models.CheckAuth(username, password)
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !exist {
		g.Response(http.StatusOK, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"token": token,
	})
	/*if ok, _ := valid.Valid(&a); ok {
		if ok := models.CheckAuth(username, password); ok {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			//logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})*/
}
