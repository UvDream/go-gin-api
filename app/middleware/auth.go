package middleware

import (
	"errors"
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

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	Signkey          string = "UvDream"
)

// AuthRequired 验证是否登陆以及是否过期
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		utilGin := util.Gin{Ctx: c}
		// 无token
		if token == "" {
			utilGin.Response(200, "fail", "请先登陆")
			c.Abort()
			return
		}
		// 有token
		j := NewJwt()
		msg, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				utilGin.Response(40001, "fail", "授权过期")
				c.Abort()
				return
			}
			utilGin.Response(40001, "fail", "请重新登陆")
			c.Abort()
			return
		}
		c.Set("msg", msg)

	}
}

// NewJwt 新建jwt实例
func NewJwt() *Jwt {
	return &Jwt{
		SigningKey: []byte(GetSignKey()),
	}
}

// GetSignKey 获取SignKey
func GetSignKey() string {
	return Signkey
}

// GenerateToken 生成令牌
func GenerateToken(c *gin.Context, u model.User) (string, error) {
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
	return token, err
}

// CreateToken 生成token
func (j *Jwt) CreateToken(msg MoreMessage) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, msg)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析token
func (j *Jwt) ParseToken(tokenString string) (*MoreMessage, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MoreMessage{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*MoreMessage); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
