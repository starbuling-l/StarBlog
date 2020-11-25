package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/starbuling-l/StarBlog/middleware/jwt"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"github.com/starbuling-l/StarBlog/routers/api"
	v1 "github.com/starbuling-l/StarBlog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)
	r.GET("auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签
		apiv1.GET("/tags", v1.GetTags)
		//添加指定标签
		apiv1.POST("/tags", v1.AddTag)
		//修改指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
	}
	{
		//获取标签
		apiv1.GET("/articles", v1.GetArticles)
		//获取标签
		apiv1.GET("/articles/:id", v1.GetArticle)
		//添加指定标签
		apiv1.POST("/articles", v1.AddArticle)
		//修改指定标签
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定标签
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
