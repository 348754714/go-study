package main

import "fmt"

// 指针
func main() {
	// 1. &：取地址
	// 2. *：根据地址取值
	//n := 8
	//fmt.Println(&n)
	//fmt.Println(*&n)
	//fmt.Printf("%T\n", &n)
	//fmt.Printf("%T\n", *&n)

	//var a *int // nil
	//*a = 100   // 无内存地址无法赋值
	//fmt.Println(*a)
	// new函数申请一个内存地址
	var a = new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
	fmt.Println(a)

	// new很少用，给基本数据类型申请内存地址，string，int... 返回类型指针 *string *int
	// make 也是申请分配内存地址，作用于silce，map，chan（通道），返回对应类型本身
}
