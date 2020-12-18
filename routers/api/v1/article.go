package v1

import (
	"github.com/starbuling-l/StarBlog/server/tag_service"
	"net/http"

	"github.com/starbuling-l/StarBlog/pkg/app"
	"github.com/starbuling-l/StarBlog/pkg/e"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"github.com/starbuling-l/StarBlog/pkg/util"
	"github.com/starbuling-l/StarBlog/server/article_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary 获取单个文章
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {
	g := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		g.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, article)

	/*	code := e.INVALID_PARAMS
		var data interface{}
		if !valid.HasErrors() {
			if models.ExistArticleByID(id) {
				data = models.GetArticle(id)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_ARTICLE
			}
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})*/
}

// @Summary 获取多个文章
// @Produce  json
// @Param tag_id query int true "tag_id"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/article [get]
func GetArticles(c *gin.Context) {
	g := app.Gin{C: c}

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = state
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	articleService := article_service.Article{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	lists, err := articleService.GetAll()
	if err != nil {
		g.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	total, err := articleService.Count()
	if err != nil {
		g.Response(http.StatusInternalServerError, e.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	data["lists"] = lists
	data["total"] = total

	g.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary 添加文章
// @Produce  json
// @Param tag_id query int true "tag_id"
// @Param title query string true "title"
// @Param desc query string true "desc"
// @Param content query string true "content"
// @Param state query int false "State"
// @Param created_by query string true "created_by"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
	g := app.Gin{C: c}

	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	coverImageUrl := c.PostForm("cover_image_url")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标题ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		g.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	tag := tag_service.Tag{ID: tagId}

	exits, err := tag.ExitByID()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exits {
		g.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
	}

	article := article_service.Article{
		TagID:         tagId,
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
		CreatedBy:     createdBy,
	}

	err = article.Add()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, nil)

	/*	code := e.INVALID_PARAMS
		if !valid.HasErrors() {
			if models.ExistTagByID(tagId) {

				code = e.SUCCESS

				data := make(map[string]interface{})
				data["tag_id"] = tagId
				data["title"] = title
				data["desc"] = desc
				data["content"] = content
				data["created_by"] = createdBy
				data["state"] = state
				models.AddArticle(data)
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})*/
}

// @Summary 编辑文章
// @Produce  json
// @Param id path int true "id"
// @Param tag_id path int true "tag_id"
// @Param title query string true "title"
// @Param desc query string true "desc"
// @Param content query string true "content"
// @Param modified_by query string true "modified_by"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [put]
func EditArticle(c *gin.Context) {
	g := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")
	coverImageUrl := c.PostForm("cover_image_url")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标题ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题不能为空")
	valid.MaxSize(desc, 255, "desc").Message("简述不能为空")
	valid.MaxSize(content, 321443214, "content").Message("内容不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100个字符")

	if valid.HasErrors() {
		g.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	article := article_service.Article{
		ID:            id,
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		ModifiedBy:    modifiedBy,
	}
	existArticle, err := article.ExistByID()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !existArticle {
		g.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tag := tag_service.Tag{ID: tagId}
	existTag, err := tag.ExitByID()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !existTag {
		g.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = article.Edit()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_EDIT_ARTICLE_FAIL, nil)
	}

	g.Response(http.StatusOK, e.SUCCESS, nil)
	/*	code := e.INVALID_PARAMS
		if !valid.HasErrors() {
			if models.ExistArticleByID(id) {
				if models.ExistTagByID(tagId) {
					code = e.SUCCESS

					data := make(map[string]interface{})
					if tagId > 0 {
						data["tag_id"] = tagId
					}
					if title != "" {
						data["title"] = title
					}
					if desc != "" {
						data["desc"] = desc
					}
					if content != "" {
						data["content"] = content
					}
					data["modified_by"] = modifiedBy
					models.EditArticle(id, data)
				} else {
					code = e.ERROR_NOT_EXIST_TAG
				}
			} else {
				code = e.ERROR_NOT_EXIST_ARTICLE
			}
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})*/
}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	g := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		g.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	article := article_service.Article{ID: id}
	exist, err := article.ExistByID()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exist {
		g.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
	}

	err = article.Delete()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, nil)
/*	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})*/
}
