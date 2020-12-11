package main

import "fmt"

//布尔值
//bool类型声名布尔值
//bool值默认为false
//go不允许将整型强制转换城bool
//布尔型无法参与数值运算，无法转换成其他数据类型
func main() {
	b1 := true
	var b2 bool
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2) // %v 打印变量值
}
