package route

import (
	"go-gin-api/app/api"
	"go-gin-api/app/middleware"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	// 测试路由
	r.GET("/ping", api.Ping)
	// 404
	r.NoRoute(func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	// 用户模块
	userRouter := r.Group("/user")
	{
		// 注册接口
		userRouter.POST("/register", api.UserRegister)
		// 登陆接口
		userRouter.POST("/login", api.UserLogin)
		auth := userRouter.Group("")
		auth.Use(middleware.AuthRequired())
		{
			//退出接口
			auth.POST("/loginOut", api.UserLogout)
			// 获取用户信息
			auth.POST("/getUserInfo", api.GetUserInfo)
		}

	}
	return r
}
