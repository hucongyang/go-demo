package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/hucongyang/go-demo/conf"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * conf.Config().App.PageSize
	}
	return result
}