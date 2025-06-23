package entity

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
}

func (p *Post) Create(db *gorm.DB) error {
	return db.Create(p).Error
}
func (p *Post) Update(db *gorm.DB) error {
	return db.Save(p).Error
}
func (p *Post) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}
func GetPostById(id uint, db *gorm.DB) (*Post, error) {
	var post Post
	if err := db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func ListPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
