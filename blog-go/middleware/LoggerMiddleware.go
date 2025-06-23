package middleware

import (
	"blog/common"
	"blog/entity"
	"blog/utils"
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// LoggerMiddleware 是一个简单的日志中间件，记录每个请求的开始时间和结束时间
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.GetUserByToken(c)
		// 获取请求开始的时间
		startTime := time.Now()
		w := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = w

		// 执行后续的中间件或路由处理函数
		c.Next()

		// 记录请求处理完成后的时间
		endTime := time.Now()

		// 计算请求处理耗时
		latency := endTime.Sub(startTime)
		latencyString := latency.String()
		db := common.GetDB()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "Failed to read request body"})
			return
		}
		responseBody := w.body.String()
		var userId uint = 0
		if user != nil {
			userId = user.ID
		}
		if err := db.Create(&entity.Log{UserID: uint(userId), InParams: string(body), OutParameter: responseBody, Time: latencyString, Url: c.Request.URL.Path}).Error; err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{"error": "Failed to create log"})
		}
	}
}
