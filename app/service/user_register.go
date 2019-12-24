package service

import (
	"fmt"
	"go-gin-api/app/model"
	"go-gin-api/app/util"
	"time"
)

// UserRegisterService 注册接口参数
type UserRegisterService struct {
	UserName        string    `json:"user_name" `
	NickName        string    `json:"nick_name"`
	Sex             string    `json:"sex"`
	Birthday        time.Time `json:"birthday"`
	Phone           string    `json:"phone"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	Avatar          string    `json:"avatar"`
	Address         string    `json:"address"`
}

// Register 注册用户
func (s *UserRegisterService) Register() *util.Response {
	fmt.Println(s)
	user := model.User{
		UserName: s.UserName,
		NickName: s.NickName,
		Sex:      s.Sex,
		Phone:    s.Phone,
		Avatar:   s.Avatar,
		Address:  s.Address,
		Birthday: s.Birthday,
		Status:   model.Active,
	}

	// 表单校验
	if err := s.valid(); err != nil {
		return err
	}

	// 加密密码
	
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return &util.Response{
			Code:    40000,
			Message: "注册失败",
		}
	}

	return &util.Response{
		Code:    200,
		Message: "注册成功",
	}
}

// 参数校验
func (s *UserRegisterService) valid() *util.Response {
	if s.Password != s.PasswordConfirm {
		return &util.Response{
			Code:    4001,
			Message: "两次输入的密码不一致",
		}
	}
	count := 0
	// 昵称校验
	model.DB.Model(&model.User{}).Where("nick_name=?", s.NickName).Count(&count)
	if count > 0 {
		return &util.Response{
			Code:    40002,
			Message: "昵称已被注册",
		}
	}
	// 手机号校验
	count = 0
	model.DB.Model(&model.User{}).Where("phone=?", s.Phone).Count(&count)
	if count > 0 {
		return &util.Response{
			Code:    4002,
			Message: "手机号已被注册",
		}
	}
	count = 0
	model.DB.Model(&model.User{}).Where("user_name=?", s.UserName).Count(&count)
	if count > 0 {
		return &util.Response{
			Code:    4002,
			Message: "用户名已被注册",
		}
	}
	return nil
}
