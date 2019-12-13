package product

import (
	"fmt"
	parambind "go-gin-api/app/controller/param_bind"
	paramverify "go-gin-api/app/controller/param_verify"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// Add 新增
func Add(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}

	// 参数绑定
	s, e := util.Bind(&parambind.ProductAdd{}, c)
	if e != nil {
		utilGin.Response(-1, e.Error(), nil)
		return
	}

	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	_ = validate.RegisterValidation("NameValid", paramverify.NameValid)

	if err := validate.Struct(s); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	// 业务处理...

	utilGin.Response(1, "success", nil)
}

// Edit 编辑
func Edit(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}

// Delete 删除
func Delete(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}

// Detail 详情
func Detail(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}
