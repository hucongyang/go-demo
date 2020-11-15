package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-demo/pkg/errorCode"
)

type Gin struct {
	C *gin.Context
}

// 封装response返回的结构，便于以后修改
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  errorCode.GetMessage(errCode),
		"data": data,
	})
	return
}
