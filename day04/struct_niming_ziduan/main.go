package main

import "fmt"

// 匿名字段

type student struct {
	string
	int
}

func main() {
	p1 := student{
		"name1", 18,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
