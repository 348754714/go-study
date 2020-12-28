package main

import "fmt"

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p1 person
	p1.name = "旺旺"
	p1.age = 20
	p1.gender = "男"
	p1.hobby = []string{"篮球", "足球", "台球"}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)

	// 匿名结构体,多用于临时场景
	var s struct { // s 是一个匿名结构体类型变量
		name string
		age  int
	}
	s.name = "fdshjaf"
	s.age = 100
	fmt.Println(s)
}
