package main

import (
	"fmt"
)

// 常量,定义后在程序运行期间不会改变，不能再次赋值
const pi = 3.14159265357

const (
	statusOk = 200
	notFound = 404
)

// 后面不写值的默认跟上面一行值一样
const (
	n1 = 100
	n2
	n3
)

// iota 在const中第一次出现重置为0,每新增一行常量声名iota+1
const (
	i1 = iota // 0
	i2        // 1
	i3        // 2
	i4        // 3
)

const (
	b1 = iota // 0
	b2        // 1
	_         // 2
	b3        // 3
)

// iota 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        // 3
)

// 多个常量在一行声名
const (
	d1, d2 = iota + 1, iota + 2 // 1,2

	d3, d4 = iota + 1, iota + 2 // 2,3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 左移10位
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("pi:", pi)

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("i1:", i1)
	fmt.Println("i2:", i2)
	fmt.Println("i3:", i3)
	fmt.Println("i4:", i4)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	//fmt.Println("TB:",TB) // 太大无法打印
	//fmt.Println("PB:",PB) // 太大无法打印
}
