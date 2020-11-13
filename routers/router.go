package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-demo/conf"
)

// 路由文件

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(conf.Config().RunMode)

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test1",
		})
	})
	return router
}