/*
 * @Author: wangzhongjie
 * @Date: 2019-12-13 09:45:41
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2019-12-13 09:50:05
 * @Description:请求返回格式化
 * @Email: UvDream@163.com
 */
package util

import (
	"github.com/gin-gonic/gin"
)

// Gin 结构体
type Gin struct {
	Ctx *gin.Context
}
type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response 格式化返回格式
func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}
