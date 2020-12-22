package main

import "fmt"

// make()构造切片
// 切片就是一个框，框住了一块连续的内存，属于引用类型，真正的数据保存在数组中
func main() {
	s1 := make([]int, 5, 10)
	fmt.Println(len(s1), cap(s1))

	//var s1 []int         // len=0 cap=0 s1==nil
	//s1 := []int{}        // len=0,cap=0 s1!=nil
	//s1 := make([]int, 0) // len=0,cap=0 s1!=nil
	// 判断切片是否为空 用len(s1) == 0 判断
}
