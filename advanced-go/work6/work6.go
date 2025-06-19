package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) printInfo() {
	fmt.Printf("Employee ID: %d, Name: %s, Age: %d", e.EmployeeID, e.Name, e.Age)
}

func main() {

	e := Employee{Person{"John", 30}, 1234}
	e.printInfo()
}
