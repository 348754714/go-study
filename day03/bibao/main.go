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

// 定义一个函数对f2进行包装 无返回值f1不能接收
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

// f1中想调用f2
// 定义一个函数对f2进行包装，返回值为func()，f1可以接收
func f4(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}
func main() {
	f3(f2, 100, 200)
	ret := f4(f2, 200, 300) // 将原来需要传递两个int的函数返回成了一个不需要传参的函数
	f1(ret)
}
