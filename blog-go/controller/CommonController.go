package controller

import (
	"blog/entity"
	"blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment entity.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	service.CreateComment(c, comment)
}

func ListComments(c *gin.Context) {
	comments, err := service.ListComments(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments, "message": "Comments retrieved successfully"})
}
