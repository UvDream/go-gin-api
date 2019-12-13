package route

import (
	"go-gin-api/app/controller/product"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(engine *gin.Engine) {
	// 测试路由
	engine.GET("/ping", func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil)
	})
	// 404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	ProductRouter := engine.Group("/product")
	{
		// 新增产品
		ProductRouter.POST("", product.Add)

		// 更新产品
		ProductRouter.PUT("/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/:id", product.Detail)
	}
}
