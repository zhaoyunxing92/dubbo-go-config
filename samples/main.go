package main

import (
	"fmt"
	"zhaoyunxing92/dubbo-go-config/config"
)

func main() {
	fmt.Println("start load config file")
	conf := config.Load()
	application, _ := conf.GetApplicationConfig()
	fmt.Println("application name:", application.Name)
}
