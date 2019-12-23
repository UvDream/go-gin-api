package api

import (
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(200, "success", "注册成功!")
}

// UserLogin 用户登陆
func UserLogin(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(200, "success", "登陆成功!")
}

// UserLogout 用户登陆
func UserLogout(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(200, "success", "退出成功!")
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(200, "success", "查询个人信息!")
}
