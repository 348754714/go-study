package main

import (
	"fmt"
	"strings"
)

// 闭包是什么
// 闭包是一个函数，包含了他外部作用域的一个变量
// 底层原理
// 1. 函数可以作为一个返回值
// 2. 函数内部查找变量顺序，先在自己内部找，找不到往外层找
// 3. 闭包=函数 + 外部变量的引用
func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder1(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	ret := adder()   // ret为adder的返回值 func(int) int
	ret2 := ret(200) // func(int) int 接收参数y=200 执行 x += y, x在adder中向外层找100
	fmt.Println(ret2)

	ret3 := adder1(100)
	ret4 := ret3(200)
	fmt.Println(ret4)

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunx := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("test1.jpg"))
	fmt.Println(txtFunx("test"))

	f1, f2 := calc(10)
	// 每一次执行，base都发生了变化
	fmt.Println(f1(1), f2(2)) // 11 9
	fmt.Println(f1(3), f2(4)) // 12 8
	fmt.Println(f1(5), f2(6)) // 13 7

}
