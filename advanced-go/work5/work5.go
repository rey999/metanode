package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() int
	Perimeter() int
}

// 实现接口的具体结构体
type Rectangle struct {
	Width, Height int
}

// 实现 Area 方法
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// 实现 Perimeter 方法
func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体表示圆形
type Circle struct {
	Radius int
}

// Area 实现计算面积
func (c Circle) Area() int {
	return int(math.Pi * float64(c.Radius) * float64(c.Radius))
}

// Perimeter 实现计算周长
func (c Circle) Perimeter() int {
	return int(2 * math.Pi * float64(c.Radius))
}
func main() {
	var s Shape          // 声明一个接口变量
	r := Rectangle{3, 4} // 创建结构体实例
	s = r                // 将结构体赋值给接口
	fmt.Println("Area:", s.Area())

	var s1 Shape            // 声明一个接口变量
	c1 := Circle{Radius: 5} // 创建圆形实例
	s1 = c1                 // 将圆形赋值给接口
	fmt.Println("Area:", s1.Area())
	fmt.Println("Perimeter:", s1.Perimeter())
}
