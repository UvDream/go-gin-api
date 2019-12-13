package parambind

import (
	paramverify "go-gin-api/app/controller/param_verify"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

// ProductAdd 接受参数
type ProductAdd struct {
	Name string `form:"name" json:"name" binding:"required,NameValid"`
}

func init() {
	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", paramverify.NameValid)
	}
}
