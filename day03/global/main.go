package main

import "fmt"

// 变量作用域
func f1() {
	// 变量查找顺寻
	//1 在函数内部找
	//x := 1 // 内部变量无法在函数外部使用
	//2 往函数外部查找，一直找到全局
	//3 都找不到报错
	fmt.Println(x)
}

var x = 100

func main() {
	f1()
}
