package main

import "fmt"

// 切片slice
func main() {
	// 定义
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"as", "fdsyaf", "fyusf"}
	fmt.Println(s1, s2)
	fmt.Println(len(s1), cap(s1), len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s3 := a1[0:4]           // 左闭右开,左包含，右不包含
	fmt.Println(s3)         // [1 2 3 4]
	s4 := a1[1:6]           // 左闭右开,左包含，右不包含
	fmt.Println(s4)         // [2 3 4 5 6]
	s5 := a1[:4]            // a1:[0:4]  [1 2 3 4]
	s6 := a1[3:]            // a1:[3:len(a1)] [4 5 6 7 8 9 10]
	s7 := a1[:]             // a1:[0:len(a1)] [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(s5, s6, s7) // [1 2 3 4]
	// 切片容量是底层数组从切片第一个元素到最后元素的数量
	fmt.Println(len(s4), cap(s4)) // 5 9
	fmt.Println(len(s5), cap(s5)) // 4 10
	fmt.Println(len(s6), cap(s6)) // 7 7
	// 切片再切片
	s8 := s6[3:]
	fmt.Println(len(s8), cap(s8)) // 4 4
	s6[0] = 1                     // 切片为引用类型修改切片的值，底层数组的值也变化
	fmt.Println(a1)
	fmt.Println(s6)
	a1[6] = 1300
	fmt.Println(s6) // 切片为引用类型修改底层数组的值，切片的值也变化
	fmt.Println(s8)

}
