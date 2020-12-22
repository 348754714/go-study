package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a, len(a), cap(a))
	sort.Ints(a) // srot排序
	fmt.Println(a)
}
