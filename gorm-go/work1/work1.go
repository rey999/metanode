package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Students struct {
	Id    int64  `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Age   int    `gorm:"column:age" json:"age"`
	Grade string `gorm:"column:grade" json:"grade"`
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别（Silent、Info、Warn、Error、Slow）
			Colorful:      true,        // 是否彩色输出
		},
	)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	db.AutoMigrate(&Students{})

	stu := Students{Name: "张三", Age: 20, Grade: "三年级"}

	db.Create(&stu)

	fmt.Println(stu)

	var stu1 Students = Students{}
	db.Where("age > 18").First(&stu1)
	fmt.Println(stu1)

	stu1.Grade = "四年级"
	db.Save(&stu1)

	var stu2 Students = Students{}
	db.Where("age > 18").First(&stu2)
	fmt.Println(stu2)

	db.Where("age < 15").Delete(&Students{})

}
