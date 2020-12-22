package main

import "fmt"

// goto
func main() {
	// 常用标识位方式控制外层for跳出
	//var flag = false
	//for i := 0; i < 10; i++ {
	//	for j := 'A'; j < 'Z'; j++ {
	//		if j == 'C' {
	//			flag = true
	//			break
	//		}
	//		fmt.Printf("%v-%c\n", i, j)
	//	}
	//	if flag {
	//		break
	//	}
	//}

	//goto + lable 跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX // 跳到指定的标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
XX: // lable标签
	fmt.Println("over")
}
