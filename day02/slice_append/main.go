package main

import "fmt"

// append slice
func main() {
	s1 := []string{"dsaf", "fdhsjaf"}
	fmt.Println(len(s1), cap(s1)) // 2 2

	// 必须用原理的变量接受append的返回值
	// append放不下时，go底层会把底层数组换一个
	s1 = append(s1, "fdsafdsafds")
	fmt.Println(len(s1), cap(s1)) //3 4
	fmt.Print(s1)
	s1 = append(s1, "fd", "fdsa", "fdsffff")
	fmt.Println(len(s1), cap(s1)) //6 8
	ss := []string{"wuhan", "xian", "chengdu"}
	// ss... append时将ss拆开
	s1 = append(s1, ss...)
	fmt.Println(len(s1), cap(s1)) //9 16
}
