/*
 * @Author: wangzhongjie
 * @Date: 2019-12-13 09:10:41
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2019-12-13 11:47:59
 * @Description:路由文件
 * @Email: UvDream@163.com
 */
package route

import (
	"go-gin-api/app/controller/product"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c} //初始化操作
		utilGin.Response(200, "这是返回信息", nil)
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
	})
	ProductRouter := engine.Group("")
	{
		// 新增产品
		ProductRouter.POST("/product", product.Add)

		// 更新产品
		ProductRouter.PUT("/product/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/product/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/product/:id", product.Detail)
	}
}
