package main

import "fmt"

// for循环

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	////变种1
	//var i1 = 5
	//for ; i1 < 10; i1++ {
	//	fmt.Println(i1)
	//}
	////变种2
	//var i2 = 5
	//for i2 < 10 {
	//	fmt.Println(i2)
	//	i2++
	//}

	// 无限循环
	//for {
	//	fmt.Println("123")
	//}

	//for rang
	//s := "f反对和沙发"
	//for i, v := range s {
	//	fmt.Printf("%d %c\n", i, v)
	//}

	// 99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			c := i * j
			fmt.Printf("%d x %d = %d\t", i, j, c)
		}
		fmt.Println()
	}
}
