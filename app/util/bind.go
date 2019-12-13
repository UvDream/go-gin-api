package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Bind 绑定参数
func Bind(s interface{}, c *gin.Context) (interface{}, error) {
	b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(s, b); err != nil {
		return nil, err
	}
	return s, nil
}
