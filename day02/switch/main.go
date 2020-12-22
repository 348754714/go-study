package main

import "fmt"

// switch 简化判断
func main() {
	//var n = 5
	switch n := 5; n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("default")
	}

	switch n1 := 1; n1 {
	case 1, 2, 3, 4, 5:
	case 6, 7, 8, 9, 10:
	default:
	}

	switch n2 := 1; n2 {
	case 1, 2, 3, 4, 5:
		fmt.Println("奇数")
	case 6, 7, 8, 9, 10:
		fmt.Println("偶数")
	default:
	}

	var n3 = 8
	switch {
	case n3 > 10:
		fmt.Println(n3)
	case n3 < 10:
		fmt.Println(n3)
	default:
		fmt.Println("hhhh")
	}

	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough // case下穿
	case s == "b":
		fmt.Println("b") // 输出ab
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

}
