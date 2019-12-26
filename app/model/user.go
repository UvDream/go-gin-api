package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

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
	// PassWordLevel 加密等级
	PassWordLevel = 12
)

// SetPassword 加密密码
func (u *User) SetPassword(password string) error {
	fmt.Println(password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordLevel)
	if err != nil {
		return err
	}
	u.PWD = string(bytes)
	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PWD), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
