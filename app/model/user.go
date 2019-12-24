package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户类
type User struct {
	gorm.Model
	ID       string `json:"user_id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Sex      string `json:"sex"`
	Birthday string `json:"birthday"`
	Phone    string `json:"phone"`
	PWD      string `json:"password"`
}

// LoginReq 登陆类
type LoginReq struct {
	Phone string `json:"phone"`
	Pwd   string `json:"password"`
}
