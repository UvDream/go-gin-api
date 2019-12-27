package controller

import (
	"fmt"
	"go-gin-api/app/middleware"
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
	var l service.UserLoginService
	// 获取参数
	if c.BindJSON(&l) == nil {
		res := l.Login(c)
		c.JSON(200, res)
	} else {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(200, "success", "登陆失败!")
	}

}

// UserLogout 用户退出
func UserLogout(c *gin.Context) {
	err := c.MustGet("msg").(*middleware.MoreMessage)
	// 获取用户信息
	a, _ := c.Get("msg")
	value, ok := a.(*middleware.MoreMessage)
	if ok {
		fmt.Println("打印", value.Phone)
	}
	if err != nil {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(200, "success", "Token有效!")
	}

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
