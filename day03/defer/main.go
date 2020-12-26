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
	return x // 1返回值赋值=5 2 defer x++ 3 ret x=5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值赋值x = 5 2defer x++ ret x = 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值 = y = x = 5, x++, ret = y = 5
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是x的副本
	}(x)
	return 5 // 返回值 = x = 5， defer 中x为接收x的副本然后++，ret=x=5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 返回值 = x =5， defer中接收x的副本返回int无人接收 ，ret = x = 6
}

// 传一个x的指针到匿名函数
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 返回值=x=5, defer接收x的地址，x指针++，ret = x = 6
}

//defer面试题
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	//deferDemo()
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
	fmt.Println(f6()) // 6

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // defer只能挂起外层的func，内部的calc("10", a, b)会当时执行
	a = 0
	defer calc("2", a, calc("20", a, b)) // a=0,b=1 , calc("20", a, b) = "20",0,1,1 = 1;calc("2", a, calc("20", a, b)) = "2",0,1,1
	b = 1
	// calc("10", 1, 2) -> 10 1 2 3
	// defer calc("1", 1, 3))
	// calc("20", 0, 2) -> 20 0 2 2
	// defer calc("2", 0, 2)
	// calc("2", 0, 2) -> 2 0 2 2
	// calc("1", 1, 3) -> 1 1 3 4
	// 10 1 2 3
	// 20 0 2 2
	// 2 0 2 2
	// 1 1 3 4
}
