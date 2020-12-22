package main

import "fmt"

// 运算符
func main() {
	var (
		a = 5
		b = 2
	)
	//算数运算
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // 单独语句 不能放在=的右边赋值
	b-- // 单独语句
	fmt.Println(a)
	fmt.Println(b)
	// 关系运算
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	// 逻辑运算
	fmt.Println(a >= b && a != b)
	fmt.Println(a >= b || a != b)
	fmt.Println(!(a >= b && a != b))
	// 位运算：针对2进制
	// 5：101
	// 2：010

	// &位与 对应位都为1才为1
	fmt.Println(5 & 2) // 000 0
	// |位或 对应位有1个1就为1
	fmt.Println(5 | 2) // 111 7
	// ^位异或(两位不一样为1)
	fmt.Println(5 ^ 2) // 111 7
	// << 二进制左移
	fmt.Println(5 << 1) // 1010 10
	// >> 二进制右移
	fmt.Println(5 >> 1) // 10 2

	var m = int8(1)
	fmt.Println(m << 10) // 1000000000 int8 装8位00000000 -> 0

	//+=, -=, *=, /=, %= <<= >>= &= |= ^=
}
