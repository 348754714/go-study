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
	case 6, 7, 8, 9, 10:
	default:
	}

	var n3 = 10
	switch n3 {
	case n3 > 10:
		fmt.Println(n3)
	case n3 < 10:
		fmt.Println(n3)
	default:
	}

}
