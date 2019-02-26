package main

import (
	"fmt"
	"mylog/test"
)

func main() {
	cfg := LoadConfig()
	fmt.Println("config:", cfg)
	for i := 0; i < 10; i++ {
		mainLog.Infof("this is my mainlog %v",i)
	}
	// test包就相当于我们项目中的子包
	test.Test()
}
