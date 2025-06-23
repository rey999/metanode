package middleware

import (
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里编写验证逻辑
		_, err := utils.ParseToken(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		// 如果验证通过，继续处理请求
		c.Next()
	}
}
