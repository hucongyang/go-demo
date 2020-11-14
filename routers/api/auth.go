package api

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/hucongyang/go-demo/models"
	"github.com/hucongyang/go-demo/pkg/errorCode"
	"github.com/hucongyang/go-demo/pkg/logging"
	"github.com/hucongyang/go-demo/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// 根据用户名密码获取token秘钥，作为调用凭证
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := errorCode.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = errorCode.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = errorCode.SUCCESS
			}

		} else {
			code = errorCode.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	// 自定义log文件使用
	logging.Info(username, password, code, errorCode.GetMessage(code))

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorCode.GetMessage(code),
		"data": data,
	})
}
