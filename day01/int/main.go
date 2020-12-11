package main

import "fmt"

// 整形 go无法直接定义一个二进制数
//uint8 0 - 255
//uint16 0 - 65535
//unit32 0 - 4294967295
//uint64 0 - 18446744073709551615
//int8 -128 - 127
//int16 -32768-32767
//int32 -2147483648 - 2147483647
//int64 -9223372036854775808 - 9223372036854775807
//uint 32位操作系统对应uint32，64位操作系统对应uint64
//int 32位操作系统对应int32，64位操作系统对应int64
//uintptr 无符号整型，用于存放指针
func main() {
	var i1 = 10
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 2进制输出
	fmt.Printf("%o\n", i1) // 8进制输出
	fmt.Printf("%x\n", i1) // 16进制输出
	// 8进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 16禁止
	i3 := 0x7436edf
	fmt.Printf("%d\n", i3)
	fmt.Printf("%T\n", i3) // 查看变量类型

	i4 := int8(120)        // 明确指定int8类型，否则为int类型
	fmt.Printf("%T\n", i4) // 查看变量类型

	i5 := 1<<5 - 1
	fmt.Printf("%T\n", i5) // 查看变量类型
	fmt.Printf("%d\n", i5)
}
