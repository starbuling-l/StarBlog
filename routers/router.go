package routers

import (
	"net/http"

	_ "github.com/starbuling-l/StarBlog/docs"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"github.com/starbuling-l/StarBlog/pkg/upload"
	"github.com/starbuling-l/StarBlog/routers/api"
	"github.com/starbuling-l/StarBlog/routers/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.ServerSetting.RunMode)
	//获取 token
	//curl http://127.0.0.1:9000/auth?username=test&password=test123456
	r.GET("/api/auth", api.GetAuth)
	r.StaticFS("/api/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/api/upload", api.UploadImage)
	//引入swagger
	//http://127.0.0.1:9000/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//获取多个文章标签
		//curl 127.0.0.1:9000/api/v1/tags
		apiv1.GET("/tags", v1.GetTags)
		//添加指定标签
		//POST 访问http://127.0.0.1:9000/api/v1/tags?name=2&state=1&created_by=test
		apiv1.POST("/tags", v1.AddTag)
		//修改指定标签
		//PUT 访问 http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		//DELETE 访问 http://127.0.0.1:9000/api/v1/tags/1
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	{
		//获取文章
		//http://127.0.0.1:9000/api/v1/articles?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2MDY2NTg0MzgsImlzcyI6ImdpbiJ9.XWgfvD9fwyeX18NqwKDSkKOdYGan3g-eK3P5KA1u7ZI
		apiv1.GET("/articles", v1.GetArticles)
		//获取单个文章
		// GET 访问 http://127.0.0.1:9000/api/v1/articles/1
		apiv1.GET("/articles/:id", v1.GetArticle)
		//添加文章
		//POST http://127.0.0.1:9000/api/v1/articles?tag_id=3&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1
		apiv1.POST("/articles", v1.AddArticle)
		//修改指定文章
		//PUT 访问 http://127.0.0.1:9000/api/v1/articles/1?tag_id=3&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		//DELETE 访问 http://127.0.0.1:9000/api/v1/articles/1
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
