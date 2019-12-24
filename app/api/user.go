package api

import (
	"fmt"
	"go-gin-api/app/model"
	"go-gin-api/app/service"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var s service.UserRegisterService
	if c.BindJSON(&s) == nil {
		res := s.Register()
		c.JSON(200, res)
	}
}

// UserLogin 用户登陆
func UserLogin(c *gin.Context) {
	var loginReq model.LoginReq
	// 获取参数
	if c.BindJSON(&loginReq) == nil {
		fmt.Println(loginReq)
	} else {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(200, "success", "登陆失败!")
	}

}

// UserLogout 用户退出
func UserLogout(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(200, "success", "退出成功!")
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	var user model.User
	result := model.DB.Find(&user)
	p, _ := util.JSONEncode(result)
	fmt.Println(string(p))
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(200, "success", "查询个人信息!")
}
