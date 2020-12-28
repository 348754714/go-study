package main

import "fmt"

// 给任意类型添加方法
// 不能给别的包里面的类型添加方法，只能给自己的包里面的类型添加方法

type myInt int

func (i myInt) hello(v myInt) {
	fmt.Println("I'am int", v)
}
func main() {
	m := myInt(100)
	//v := int(m)
	m.hello(m)
}
