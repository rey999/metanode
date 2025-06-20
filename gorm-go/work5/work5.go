package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	User struct {
		Id        uint `gorm:"primary_key;auto_increment;column:id" json:"id"`
		Name      string
		Posts     []Post `gorm:"foreignKey:id"`
		PostCount int    `gorm:"column:post_count"`
	}

	Post struct {
		Id           uint `gorm:"primary_key;auto_increment;column:id" json:"id"`
		Title        string
		Content      string
		User         User      `gorm:"foreignKey:id"`
		Comments     []Comment `gorm:"foreignKey:id"`
		CommentCount int       `gorm:"column:comment_count"`
	}

	Comment struct {
		Id      uint `gorm:"primary_key;auto_increment;column:id" json:"id"`
		Content string
		Post    Post `gorm:"foreignKey:id"`
	}
)

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	p.User.PostCount++
	tx.Save(&p.User)
	return nil
}
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	c.Post.CommentCount++
	tx.Save(&c.Post)
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err2 := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err2 != nil {
		panic(err2)
	}
	// db.Where("1=1").Delete(&User{})
	// db.Where("1=1").Delete(&Post{})
	// db.Where("1=1").Delete(&Comment{})

	// u := User{Name: "Jinzhu"}
	// db.Create(&u)
	// db.Create(&Post{Title: "Post 1", Content: "Content 1", User: u})
	// db.Create(&Post{Title: "Post 2", Content: "Content 2", User: u})
	// db.Create(&Post{Title: "Post 3", Content: "Content 3", User: u})
	// db.Create(&Post{Title: "Post 4", Content: "Content 4", User: u})

	// db.Create(&Comment{Content: "Comment 1", Post: Post{Id: 1}})
	// db.Create(&Comment{Content: "Comment 2", Post: Post{Id: 2}})
	// db.Create(&Comment{Content: "Comment 3", Post: Post{Id: 3}})
	// db.Create(&Comment{Content: "Comment 4", Post: Post{Id: 4}})
	// db.Create(&Comment{Content: "Comment 5", Post: Post{Id: 4}})

	var posts []Post
	err1 := db.Preload("Comments").Where("1=1").Find(&posts)

	if err1.Error != nil {
		panic(err1.Error)
	}
	for _, v := range posts {
		println(v.Title, v.Content)
		for _, v1 := range v.Comments {
			println(v1.Content)
		}
	}

	var post Post
	db.Raw(
		`
		select * from posts p
		where comment_count = (
			select max(comment_count) from posts
		)`,
	).Scan(&post)

	fmt.Println(post.Title, post.Content, post.CommentCount)
}
