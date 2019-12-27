package route

import (
	"go-gin-api/app/controller"
	"go-gin-api/app/middleware"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	// 测试路由
	r.GET("/ping", controller.Ping)
	// 404
	r.NoRoute(func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	// 用户模块
	userRouter := r.Group("/user")
	{
		// 注册接口
		userRouter.POST("/register", controller.UserRegister)
		userRouter.GET("/getUserList", controller.GetUserInfo)
		// 登陆接口
		userRouter.POST("/login", controller.UserLogin)
		auth := userRouter.Group("")
		auth.Use(middleware.AuthRequired())
		{
			//退出接口
			auth.POST("/loginOut", controller.UserLogout)
			// 获取用户信息
			auth.GET("/getUserInfo", controller.GetUserInfo)
		}

	}
	return r
}
