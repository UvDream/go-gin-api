package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// DataBase 连接数据库
func DataBase(dbURL string) {
	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		fmt.Printf("连接数据库失败,err:%v\n", err)
		return
	}
	// 设置连接池
	// 空闲
	db.DB().SetMaxIdleConns(50)
	// 打开
	db.DB().SetMaxOpenConns(100)
	// 超时时间
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
}
