package main

import "fmt"

// 匿名函数
func main() {
	// 函数内部不能声名带名字的函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 如果只是调用一次，可以简写成立即执行
	func() {
		fmt.Println("heh")
	}()

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
