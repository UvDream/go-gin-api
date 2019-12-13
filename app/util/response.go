package util

import "github.com/gin-gonic/gin"

// Gin 上下文结构体
type Gin struct {
	Ctx *gin.Context
}

// Response 返回结构体定义
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Response 返回信息处理
func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}
