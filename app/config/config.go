package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config 配置结构体
type Config struct {
	// 数据库
	MysqlDNS string `yaml:"MYSQL_DSN"`
	// redis
	RedisADDR string `yaml:"REDIS_ADDR"`
	// redis密码
	RedisPW string `yaml:"REDIS_PW"`
	// redis数据库
	RedisDB string `yaml:"REDIS_DB"`
	// redis 秘钥
	SessionSecret string `yaml:"SESSION_SECRET"`
	// 模式
	GinMode string `yaml:"GIN_MODE"`
	// 日志级别
	LogLevel string `yaml:"LOG_LEVEL"`
	// 端口号
	AppPort string `yaml:"APP_PORT"`
}

// GetConfig 读取配置
func (c *Config) GetConfig() *Config {
	path, err := filepath.Abs("config.yaml")
	if err != nil {
		fmt.Printf("读取配置失败,err:%v\n", err)
		return nil
	}
	fmt.Println("获取到配置文件:", path)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}
