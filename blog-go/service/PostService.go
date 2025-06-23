package service

import (
	"blog/entity"
	"blog/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context, post entity.Post) error {

	// 获取当前用户
	user, err := utils.GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return err
	}
	post.UserID = user.ID

	// 保存文章到数据库
	if err := post.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return err
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
	return nil
}

func GetPostById(c *gin.Context, id uint) (error, post *entity.Post) {
	post, err := entity.GetPostById(id, db)
	if err != nil {
		return error, nil
	}
	c.JSON(http.StatusOK, gin.H{"data": post, "message": "Post retrieved successfully"})
	return nil, error
}
func UpdatePost(c *gin.Context, post entity.Post) error {
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return fmt.Errorf("id is required")
	}
	// 获取当前用户
	user, err := utils.GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return err
	}
	storagePost, err := entity.GetPostById(post.ID, db)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return err
	}

	if user.ID != storagePost.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this post"})
		return err
	}

	if err := post.Update(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return err
	}
	return nil
}

func DeletePost(c *gin.Context) error {
	user, _ := utils.GetUserByToken(c)
	id := c.Param("id")
	var post entity.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return err
	}
	if post.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return fmt.Errorf("Forbidden")
	}
	if err := post.Delete(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return err
	}
	return nil
}

func ListPosts(c *gin.Context) (error, []entity.Post) {
	posts, err := entity.ListPosts(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list posts"})
		return err, nil
	}
	c.JSON(http.StatusOK, gin.H{"data": posts, "message": "Posts retrieved successfully"})
	return nil, posts
}
