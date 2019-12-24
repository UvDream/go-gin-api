package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User 用户类
type User struct {
	gorm.Model
	UserName string    `json:"user_name"`
	NickName string    `json:"nick_name"`
	Sex      string    `json:"sex"`
	Birthday time.Time `json:"birthday"`
	Phone    string    `json:"phone"`
	// column:beast_id 重写列
	PWD     string `json:"password" gorm:"column:password"`
	Avatar  string `json:"avatar"`
	Status  string `json:"status"`
	Address string `json:"address"`
}

const (
	// Active 激活用户
	Active string = "active"
	// InActive 未激活用户
	InActive string = "inactive"
)

// LoginReq 登陆类
type LoginReq struct {
	Phone string `json:"phone"`
	Pwd   string `json:"password"`
}
