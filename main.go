package main

import (
	"go-gin-api/app/config"
	"go-gin-api/app/model"
	"go-gin-api/app/route"
	"log"
	"os"
)

func main() {
	var c config.Config
	config := *c.GetConfig()
	// 初始化配置
	c.Init()
	// 装载路由
	r := route.NewRouter()
	// 启动服务
	if err := r.Run(config.AppPort); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer model.DB.Close()

}
