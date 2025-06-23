package utils

import (
	"blog/config"
	"blog/entity"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(config.ConfigData.SecretKey) // 签名密钥，应与生成 token 时使用的密钥一致

func ParseToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := GetToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否符合预期（例如 HMAC-SHA）
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 检查 token 是否有效
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return token, nil
}

func GetToken(c *gin.Context) string {
	return c.GetHeader("Authorization")
}

func GetUserByToken(c *gin.Context) (*entity.User, error) {
	token, err := ParseToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	id, ok := claims["id"].(uint)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	user := entity.User{}
	user.ID = id
	user.username = username
	return &user, nil

}
