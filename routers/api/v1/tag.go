package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-demo/service/tag_service"
	"github.com/unknwon/com"
	"net/http"

	"github.com/hucongyang/go-demo/conf"
	"github.com/hucongyang/go-demo/models"
	"github.com/hucongyang/go-demo/pkg/app"
	"github.com/hucongyang/go-demo/pkg/errorCode"
	"github.com/hucongyang/go-demo/pkg/util"
)

// @Summary 获取文章标签
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} gin.H "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [get]
// 获取多个文章标签
func GetTags(c *gin.Context) {
	appGin := app.Gin{C: c}
	name := c.DefaultQuery("name", "")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appGin.Response(http.StatusOK, errorCode.INVALID_PARAMS, nil)
		return
	}

	tagService := tag_service.Tag{Name: name}
	exists, err := tagService.ExistByName()
	if err != nil {
		appGin.Response(http.StatusOK, errorCode.ERROR_CHECK_EXIST_TAG_FAIL, nil)
		return
	}
	if !exists {
		appGin.Response(http.StatusOK, errorCode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	tag, err := tagService.GetAll()
	if err != nil {
		appGin.Response(http.StatusOK, errorCode.ERROR_GET_TAG_FAIL, nil)
		return
	}
	appGin.Response(http.StatusOK, errorCode.SUCCESS, tag)
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} gin.H "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
// 新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := errorCode.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = errorCode.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = errorCode.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorCode.GetMessage(code),
		"data": make(map[string]string),
	})
}

// 修改文章标签
func EditTag(c *gin.Context) {

}

// 删除文章标签
func DeleteTag(c *gin.Context) {

}
