package utils

import (
	"blog/config"
	"blog/entity"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// var jwtSecret = []byte(config.ConfigData.SecretKey) // 签名密钥，应与生成 token 时使用的密钥一致

func ParseToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := GetToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ConfigData.SecretKey), nil
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
	id, ok := claims["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	user := entity.User{}
	user.ID = uint(id)
	user.Username = username
	return &user, nil

}
