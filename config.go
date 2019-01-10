package main

// 日志级别的配置应该在配置文件中
// 这里为了方便直接在文件中定义了
var DebugLevel = "trace"

type Config struct {
	DebugLevel string
}

func LoadConfig() *Config {
	initLogRotator("./test.log")
	setLogLevels(DebugLevel)
	return &Config{
		DebugLevel: DebugLevel,
	}
}
