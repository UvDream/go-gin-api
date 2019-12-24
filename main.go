package main

import (
	"go-gin-api/app/config"
	"go-gin-api/app/model"
	"go-gin-api/app/route"
	"log"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	var c config.Config
	config := *c.GetConfig()
	// 初始化配置
	c.Init()
	// 装载路由
	r := route.NewRouter()
	url := ginSwagger.URL("http://localhost:8000/docs/swagger.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// 启动服务
	if err := r.Run(config.AppPort); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer model.DB.Close()

}
