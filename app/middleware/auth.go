package middleware

import (
	"fmt"
	"go-gin-api/app/model"
	"go-gin-api/app/util"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// Jwt 签名结构
type Jwt struct {
	SigningKey []byte
}

// MoreMessage 加密更多信息
type MoreMessage struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// AuthRequired 验证是否登陆以及是否过期
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		// 无token
		if token == "" {
			utilGin := util.Gin{Ctx: c}
			utilGin.Response(200, "fail", "请先登陆")
			c.Abort()
			return
		}
		// 有token
		fmt.Println("获取的token是:", token)
	}
}

// GenerateToken 生成令牌
func GenerateToken(c *gin.Context, u model.User)(string,error) {
	fmt.Println(u)
	j := Jwt{
		SigningKey: []byte("UvDream"),
	}
	fmt.Println(j)
	msg := MoreMessage{
		ID:    u.ID,
		Name:  u.UserName,
		Phone: u.Phone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "UvDream",                       //签名的发行者
		},
	}
	token, err := j.CreateToken(msg)
	return token,err
}

// CreateToken 生成token
func (j *Jwt) CreateToken(msg MoreMessage) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, msg)
	return token.SignedString(j.SigningKey)
}
