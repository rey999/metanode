package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

func (c *Comment) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

func ListComments(db *gorm.DB) ([]Comment, error) {
	var comments []Comment
	if err := db.Preload("User").Preload("Post").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
