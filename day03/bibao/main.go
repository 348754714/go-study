package main

import "fmt"

// 闭包
func f1(f func()) {
	fmt.Println("f1 func")
	f()
}

func f2(x, y int) {
	fmt.Println("f2 func")
	fmt.Println(x + y)
}

// f1中想调用f2
// 定义一个函数对f2进行包装
func f3(f func(int, int), m, n int) {
	//tmp := func() {
	//	f(m, n)
	//}
	//tmp()
	//func() {
	//	f(m, n)
	//}()
	f(m, n)
}
func main() {
	f3(f2, 100, 200)
}
