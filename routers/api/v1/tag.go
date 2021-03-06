package v1

import (
	"net/http"

	"github.com/starbuling-l/StarBlog/pkg/app"
	"github.com/starbuling-l/StarBlog/pkg/e"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"github.com/starbuling-l/StarBlog/pkg/util"
	"github.com/starbuling-l/StarBlog/server/tag_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary 获取多个文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	g := app.Gin{C: c}
	name := c.Query("name")
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tag := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	all, err := tag.GetAll()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_GET_TAGS_FAIL, nil)
		return

	}
	count, err := tag.Count()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_COUNT_TAG_FAIL, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": all,
		"total": count,
	})

	/*	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["total"] = models.GetTagsTotal(maps)

		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": data,
		})*/
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	g := app.Gin{C: c}

	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长未100个字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		g.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	tag := tag_service.Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	}
	exist, err := tag.ExistByName()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exist {
		g.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = tag.Add()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, nil)

	/*	code := e.INVALID_PARAMS
		if !valid.HasErrors() {
			if !models.ExistTagByName(name) {
				code = e.SUCCESS
				models.AddTag(name, state, createdBy)
			} else {
				code = e.ERROR_EXIST_TAG
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})*/
}

// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	g := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modefiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.MaxSize(name, 100, "name").Message("名称最长未100个字符")
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modefiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modefiedBy, 100, "modified_by").Message("修改人最长长度为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		g.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	tag := tag_service.Tag{
		ID:         id,
		Name:       name,
		ModifiedBy: modefiedBy,
		State:      state,
	}
	exist, err := tag.ExitByID()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exist {
		g.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = tag.Edit()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, nil)

	/*	code := e.INVALID_PARAMS
		if !valid.HasErrors() {
			code = e.SUCCESS
			if models.ExistTagByID(id) {
				data := make(map[string]interface{})
				data["modified_by"] = modefiedBy
				if name != "" {
					data["name"] = name
				}
				if state != -1 {
					data["state"] = state
				}
				models.EditTag(id, data)
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})
	*/
}

// @Summary 删除文章标签
// @Produce  json
// @Param id query string true "Id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	g := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		g.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	tag := tag_service.Tag{ID: id,}
	exist, err := tag.ExitByID()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exist {
		g.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = tag.Delete()
	if err != nil {
		g.Response(http.StatusOK, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, nil)
	/*code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})
	}*/
}
