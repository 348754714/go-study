package main

import "fmt"

// fmt 占位符
func main() {
	var n = 100
	fmt.Printf("%T\n", n) // 打印类型
	fmt.Printf("%v\n", n) // 打印值
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%d\n", n) // 10进制
	fmt.Printf("%o\n", n) // 8 进制
	fmt.Printf("%x\n", n) // 16进制
	fmt.Printf("%s\n", n) // 打印字符串，整型打印字符串非想要的结果
	var s = "HELLO"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)  // HELLO
	fmt.Printf("%#v\n", s) // "HELLO"

}
