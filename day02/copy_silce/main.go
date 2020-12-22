package main

import "fmt"

// copy silce
func main() {
	a1 := []int{1, 3, 5, 7, 9}
	a2 := a1
	var a3 = make([]int, 5, 5)
	copy(a3, a1) // a3赋值为nil copy后也为nil，用make声名a3
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 切片没有删除方法
	// 将a中index=1的删掉
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1, a2, a3) // 引用导致s2变化
}
