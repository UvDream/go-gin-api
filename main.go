package main

import (
	"go-gin-api/app/route"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置gin模式
	// 1.DebugMode
	// 2.ReleaseMode
	// 3.TestMode
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	// 设置路由
	route.SetupRouter(engine)

	// 启动服务
	if err := engine.Run(":9000"); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()

}
