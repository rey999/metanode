package service

import (
	"blog/common"
	"blog/entity"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context, comment entity.Comment) error {
	db := common.GetDB()
	user, err := utils.GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return err
	}
	comment.UserID = user.ID
	if err := comment.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return err
	}
	return nil
}

func ListComments(c *gin.Context) ([]entity.Comment, error) {
	db := common.GetDB()
	comments, err := entity.ListComments(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list comments"})
		return nil, err
	}
	return comments, nil
}
