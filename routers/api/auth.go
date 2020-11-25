package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/starbuling-l/StarBlog/models"
	"github.com/starbuling-l/StarBlog/pkg/e"
	"github.com/starbuling-l/StarBlog/pkg/util"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := auth{
		Username: username,
		Password: password,
	}
	if ok, _ := valid.Valid(&a); ok {
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
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
