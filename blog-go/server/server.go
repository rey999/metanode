package server

import (
	"blog/route"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	// 创建默认的路由引擎
	r := gin.Default()
	// 定义 GET 路由
	route.InitRoutes(r)
	// 启动 HTTP 服务，默认在 0.0.0.0:8080
	r.Run(":8080")
}
