package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/starbuling-l/StarBlog/middleware/jwt"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"github.com/starbuling-l/StarBlog/routers/api"
	"github.com/starbuling-l/StarBlog/routers/api/v1"

	_ "github.com/starbuling-l/StarBlog/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)

	//引入swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签
		//http://127.0.0.1:9000/api/v1/articles?
		//token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2MDYzMTY2OTYsImlzcyI6ImdpbiJ9.PJar3SnsaPnn5O2o-_JAjoiAqkXd3ftyEmsyEPm5plo
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
