package main

import "fmt"

// 浮点数 float32 float64
func main() {
	//math.MaxFloat32 // float32最大值
	//math.MaxFloat64 // float64最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认go的小数都是float64

	f2 := float32(1.243534)
	fmt.Printf("%T\n", f2) // 默认go的小数都是float64

	//f1 = f2 // float32类型不同不能赋值float64
	//f2 = f1 // float64类型不同不能赋值float32

}
