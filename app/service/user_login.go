package service

import (
	"go-gin-api/app/middleware"
	"go-gin-api/app/model"
	"go-gin-api/app/util"
	"time"

	"github.com/gin-gonic/gin"
)

// UserLoginService 登陆类
type UserLoginService struct {
	Phone string `json:"phone"`
	Pwd   string `json:"password"`
}
type loginRes struct {
	Token    string   `json:"token"`
	UserInfo userInfo `json:"user_info"`
}
type userInfo struct {
	UserName string    `json:"user_name"`
	NickName string    `json:"nick_name"`
	Sex      string    `json:"sex"`
	Birthday time.Time `json:"birthday"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	Address  string    `json:"address"`
}

// Login 登陆服务
func (s *UserLoginService) Login(c *gin.Context) *util.Response {
	// 验证用户是否存在
	user := model.User{}
	err := model.DB.Where("phone=?", s.Phone).Find(&user).Error
	if err != nil {
		return &util.Response{
			Code:    4000,
			Message: "failed",
			Data:    "用户不存在",
		}
	}
	// 验证密码
	if err := user.CheckPassword(s.Pwd); err != nil {
		return &util.Response{
			Code:    4000,
			Message: "failed",
			Data:    "密码错误",
		}
	}
	//token
	token, err := middleware.GenerateToken(c, user)
	if err != nil {
		return &util.Response{
			Code:    40000,
			Message: "failed",
			Data:    "生成秘钥失败!",
		}
	}
	result := loginRes{
		Token: token,
		UserInfo: userInfo{
			UserName: user.UserName,
			NickName: user.NickName,
			Sex:      user.Sex,
			Birthday: user.Birthday,
			Phone:    user.Phone,
			Avatar:   user.Avatar,
			Address:  user.Address,
		},
	}
	return &util.Response{
		Code:    200,
		Message: "success",
		Data:    result,
	}
}
