package main

import "fmt"

func main() {
	fmt.Println("testing log")
	cfg := LoadConfig()
	fmt.Println("config:", cfg)
	for i := 0; i < 10; i++ {
		Mainlog.Infof("this si my log %v",i)
		TESTlog.Infof("this is test log %v",i)
	}
}
