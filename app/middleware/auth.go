package middleware

import (
	"fmt"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// AuthRequired 验证是否登陆以及是否过期
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		// 无token
		if token == "" {
			utilGin := util.Gin{Ctx: c}
			utilGin.Response(200, "fail", "请先登陆")
			c.Abort()
			return
		}
		// 有token
		fmt.Println("获取的token是:", token)
	}
}
