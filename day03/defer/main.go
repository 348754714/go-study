package main

import "fmt"

// defer
// 多用于函数结束前释放资源
// 有defer的 return x
//1 返回值 = x
//2 运行defer
//3 return命令
// 没有defer的 return x
//1 返回值 = x
//2 return命令
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("hehehe") // 延迟到函数即将返回的时候再执行，
	defer fmt.Println("hahaha") // 多个defer按照后进先出顺寻执行
	defer fmt.Println("xixixi")
	fmt.Println("end")
}

// go的函数不是原子操作
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是x的副本
	}(x)
	return 5
}
func main() {
	//deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}
