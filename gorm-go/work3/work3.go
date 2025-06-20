package main

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Employees struct {
	Id         int64  `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Salary     int64  `gorm:"column:salary" json:"salary"`
	Department string `gorm:"column:department" json:"department"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	db.AutoMigrate(&Employees{})
	db.Delete(&Employees{})

	db.Create(&Employees{Name: "张三", Salary: 10000, Department: "技术部"})
	db.Create(&Employees{Name: "李四", Salary: 20000, Department: "技术部"})
	db.Create(&Employees{Name: "王五", Salary: 30000, Department: "技术部"})
	db.Create(&Employees{Name: "赵六", Salary: 40000, Department: "运营部"})
	db.Create(&Employees{Name: "孙七", Salary: 50000, Department: "运营部"})

	sdb, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	var employees []Employees
	err1 := sdb.Select(&employees, "SELECT * FROM employees", nil)
	if err1 != nil {
		panic(err1)
	}
	for _, employee := range employees {
		println(employee.Name, employee.Salary, employee.Department)
	}

	var maxId int64
	sdb.Get(&maxId, "SELECT id FROM employees where salary = (select max(salary) from employees)", nil)
	println(maxId)

	var employee Employees
	println("工资最高的是：")
	sdb.Get(&employee, "SELECT * FROM employees where id = ?", maxId)
	println(employee.Name, employee.Salary, employee.Department)

}
