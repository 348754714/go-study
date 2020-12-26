package main

import "fmt"

// 函数
// 函数的定义
// 函数是一段代码的封装
// 函数传值都是传递的值类型

// 命名返回值就相当于声名，可以在函数中使用
func sum(x int, y int) (ret int) {
	ret = x + y
	return
}

// 可以不命名返回值
func sum2(x int, y int) int {
	ret := x + y
	return ret
}

// 没有返回值的
func sum1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {

}

// 没有参数，有返回值
func f3() int {
	return 2
}

// 多个返回值
func f5() (int, string) {
	return 1, "fdasf"
}

// 参数类型简写：参数中连续两个参数类型一直，可以将前面一个类型生了
func f6(x, y, z int, m, n string, b1, b2 bool) int {
	return x + y + z
}

// 可变长参数，参数数量可能多个
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

// GO函数没有默认参数概念
func main() {
	r := sum(10, 20)
	fmt.Println(r)

	m, s := f5()
	fmt.Println(m, s)

	f7("rainning", 1)
	f7("rainning", 1, 2, 3, 4, 5, 6, 7) // y返回类型为切片
}
