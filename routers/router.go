package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-demo/conf"
	"github.com/hucongyang/go-demo/middleware/jwt"
	"github.com/hucongyang/go-demo/routers/api"
	v1 "github.com/hucongyang/go-demo/routers/api/v1"
)

// 路由文件

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(conf.Config().RunMode)

	// 获取token凭证
	router.GET("/auth", api.GetAuth)

	apiGroupV1 := router.Group("/api/v1")
	// 使用中间件jwt
	apiGroupV1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiGroupV1.GET("/tags", v1.GetTags)
		// 新建标签
		apiGroupV1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiGroupV1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiGroupV1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return router
}
