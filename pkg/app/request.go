package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/hucongyang/go-demo/pkg/logging"
)

// 日志记录请求校验参数的log
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
}
