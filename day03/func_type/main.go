package main

import (
	"fmt"
)

// 函数类型作为返回值
// 函数是一种类型，可以作为参数，也可以作为返回值
func f1() {
	fmt.Println("shsh")
}

func f2() int {
	return 10
}

// 函数的参数可以是函数类型
func f3(f func() int) {
	ret := f()
	fmt.Println(ret)
}

func f4(x, y int) (s int, z int) {
	s = x
	z = y
	return
}

func f6(f func(int, int) int, m, n int) int {
	ret := f(m, n)
	return ret
}

func f5(f func() int) func(int, int) int {
	return ff
}

func ff(a, b int) int {
	return a + b
}
func main() {
	//a := f1
	//fmt.Printf("%T %#V\n", a, a)
	//b := f2
	//fmt.Printf("%T %#V\n", b, b)
	//f3(f2)
	////fmt.Printf("%T %#V\n", f4, f4)
	////f5(f4)
	//a := f5(f2)
	//fmt.Printf("%T %#V\n", a, a)
	//fmt.Println(&a)
	fmt.Println(f6(ff, 1, 3))

}
