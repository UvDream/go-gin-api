/*
 * @Author: wangzhongjie
 * @Date: 2019-12-13 09:10:41
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2019-12-13 09:54:25
 * @Description:路由文件
 * @Email: UvDream@163.com
 */
package route

import (
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
}
