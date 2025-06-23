package route

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Use(middleware.LoggerMiddleware())
	white := r.Group("/white")
	{
		white.POST("/register", controller.Register)
		white.POST("/login", controller.Login)
		white.POST("/post/getById", controller.GetPostById)
		white.POST("/post/list", controller.ListPosts)
		white.POST("/comment/addComment", controller.ListComments)
	}
	vip := r.Group("/vip", middleware.TokenMiddleware())
	{
		vip.POST("/post/addPost", controller.CreatePost)
		vip.POST("/post/edit", controller.UpdatePost)
		vip.POST("/post/delete", controller.DeletePost)
		vip.POST("/comment/addComment", controller.CreateComment)
	}

}
