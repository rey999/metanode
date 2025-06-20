package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Books struct {
	Id     int64   `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Title  string  `gorm:"column:title" json:"title"`
	Author string  `gorm:"column:author" json:"author"`
	Price  float64 `gorm:"column:price" json:"price"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Books{})
	db.Where("1=1").Delete(&Books{})
	db.Create(&Books{Title: "Go语言", Author: "张三", Price: 10.0})
	db.Create(&Books{Title: "Python语言", Author: "李四", Price: 200})
	db.Create(&Books{Title: "Java语言", Author: "王五", Price: 300})
	db.Create(&Books{Title: "C语言", Author: "赵六", Price: 400})
	db.Create(&Books{Title: "C++语言", Author: "孙七", Price: 500})

	sdb, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	var bkArr []Books
	err1 := sdb.Select(&bkArr, "SELECT * FROM books where price > ?", 100)
	if err1 != nil {
		panic(err1)
	}
	for _, bk := range bkArr {
		fmt.Printf("书名：%s, 作者：%s, 价格：%.2f\n", bk.Title, bk.Author, bk.Price)
	}
}
