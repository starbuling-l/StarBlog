package util

import (
	"github.com/starbuling-l/StarBlog/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//进行统一的分页
func GetPage(ctx *gin.Context) int {
	result := 0
	if page, _ := com.StrTo(ctx.Query("page")).Int(); page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
