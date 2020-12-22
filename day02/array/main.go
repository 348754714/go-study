package main

import "fmt"

// 数组 必须指定存放的类型和容量
func main() {
	var a1 [3]bool
	var a2 [4]bool
	// a1 a2长度不同就是不同类型，不能比较
	// 不初始化，都为0值（bool：false，int float：0 ,string: ""）
	fmt.Printf("a1:%T %v a2:%T %v", a1, a1, a2, a2)
	// 初始化方式1
	a1 = [3]bool{true, true, false}
	// 初始化方式2根据初始值推导长度
	a3 := [...]int{0, 1, 3}
	fmt.Println(a3)
	// 初始化方式3未赋值位补0
	a4 := [5]int{1, 2} // 补0
	fmt.Println(a4)
	// 初始化方式3 根据索引
	a5 := [5]int{0: 1, 4: 2} // 根据索引
	fmt.Println(a5)

	// 数组遍历
	citys := [3]string{"beijing", "shanghai", "wuhan"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for k, v := range citys {
		fmt.Println(k, v)
	}

	// 多维数组
	// [ [1 2] [2 3] [3 4]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{2, 3},
		[2]int{3, 4},
	}
	fmt.Println(a11)

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
