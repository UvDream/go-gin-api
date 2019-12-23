package api

import (
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// Ping 检查页面状态
func Ping(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(200, "success", nil)
}
