package test

import (
	"githun.com/naichadouban/mylog/mylog"
)
var log mylog.Logger
// UseLogger是在项目根目录下的log.go文件中调用
func UseLogger(logger mylog.Logger) {
	log = logger
}
