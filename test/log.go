package test

import (
	"github.com/naichadouban/mylog/mylog"
)

var log mylog.Logger
// UseLogger是在项目根目录下b的log.go文件中调用
func UseLogger(logger mylog.Logger) {
	log = logger
}
