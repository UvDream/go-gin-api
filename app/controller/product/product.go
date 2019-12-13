package product

import (
	parambind "go-gin-api/app/controller/param_bind"
	"go-gin-api/app/util"

	"github.com/gin-gonic/gin"
)

// Add 新增
func Add(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	if err := c.ShouldBind(&parambind.ProductAdd{}); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	utilGin.Response(1, "success", nil)
}

// Edit 编辑
func Edit(c *gin.Context) {

}

// Delete 删除
func Delete(c *gin.Context) {

}

// Detail 详情
func Detail(c *gin.Context) {

}
