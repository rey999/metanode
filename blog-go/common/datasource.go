package common

import (
	"blog/config"
	"blog/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	// SQLite 数据库文件路径（例如：test.db）
	dbPath := config.ConfigData.SqliteName
	var err error
	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移模型到数据库
	err = database.AutoMigrate(&entity.User{}, &entity.Post{}, &entity.Comment{}, &entity.Log{})
	if err != nil {
		return err
	}
	db = database

	return nil
}

func GetDB() *gorm.DB {
	return db
}
