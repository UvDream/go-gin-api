package route

import (
	"go-gin-api/app/api"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(engine *gin.Engine) {
	// 测试路由
	engine.GET("/ping", api.Ping)
	// 404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	// 用户模块
	UserRouter := engine.Group("/user")
	{
		// 注册接口
		UserRouter.POST("/register", api.UserRegister)
		// 登陆接口
		UserRouter.POST("/login", api.UserLogin)
		//退出接口
		UserRouter.POST("/loginOut", api.UserLogout)
		UserRouter.POST("/getUserInfo", api.GetUserInfo)
	}
}
