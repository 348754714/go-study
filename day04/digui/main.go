package main

import "fmt"

// 递归
// 递归要由一个明确的退出条件\问题相同，问题的规模会越来越小
func f1(n uint64) uint64 { // 阶层
	if n <= 1 {
		return n
	}
	return n * f1(n-1)
}

// 有一段楼梯台阶有n级台阶，1步，2步，共多少走法，
func taijie(n uint64) (x uint64) {
	if n <= 2 {
		return n
	}

	x = 0
	return taijie(n-1) + taijie(n-2)
}
func main() {
	//fmt.Println(f1(5))
	fmt.Println(taijie(10))
}
