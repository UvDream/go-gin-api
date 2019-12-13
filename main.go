package main

import (
	"go-gin-api/app/config"
	"go-gin-api/app/route"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置gin模式
	gin.SetMode(config.AppMode)
	engine := gin.New()
	// 记录日志
	engine.Use(gin.Logger())
	// 设置路由
	route.SetupRouter(engine)

	// 启动服务
	if err := engine.Run(config.AppPort); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
