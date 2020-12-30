package main

import "fmt"

// 二进制实际用途

// 111
// 左边1表示吃饭
// 中间表示睡觉
// 右边表示打豆豆
const (
	eat   int = 4
	sleep int = 2
	da    int = 1
)

func f(arg int) {
	fmt.Printf("%b\n", arg)
}
func main() {
	f(eat | sleep)
	f(sleep | eat)
	f(sleep | eat | da)
}
