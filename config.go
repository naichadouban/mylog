package main

import (
	"fmt"
	"goutil/appdata"
	"path/filepath"
)

// 日志级别的配置应该在配置文件中
// 这里为了方便直接在文件中定义了
var DebugLevel = "info"
var LogDir = appdata.AppDataDir("mylog", false)
var defaultLogFilename = "test.log"

type Config struct {
	LogLevel    string
	LogDir      string
	LogFilename string
}

func LoadConfig() *Config {
	// 这里正常应该是从配置文件中加载
	conf := &Config{
		LogLevel:    DebugLevel,
		LogDir:      LogDir,
		LogFilename: defaultLogFilename,
	}
	// 根据配置文件对日志系统进行配置
	fmt.Println(filepath.Join(conf.LogDir, conf.LogFilename))
	initLogRotator(filepath.Join(conf.LogDir, conf.LogFilename))
	setLogLevels(conf.LogLevel)
	return conf
}
