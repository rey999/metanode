package controller

import (
	"blog/entity"
	"blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.CreatePost(c, post)
}

func GetPostById(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, ok := json["id"].(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	err, post := service.GetPostById(c, id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post, "message": "Post retrieved successfully"})
}

func ListPosts(c *gin.Context) {
	posts, err := service.ListPosts(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts, "message": "Posts retrieved successfully"})
}

func DeletePost(c *gin.Context) {
	err := service.DeletePost(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func UpdatePost(c *gin.Context) {
	post := entity.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.UpdatePost(c, post)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}
