package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	v1 "github.com/starbuling-l/StarBlog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签
		apiv1.GET("/tags", v1.GetTags)
		//添加指定标签
		apiv1.POST("/tags", v1.AddTag)
		//修改指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
