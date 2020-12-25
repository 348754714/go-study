package main

import (
	"fmt"
	"strings"
)

// map和slice组合
func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)

	s1[0] = make(map[int]string, 1)
	s1[0][10] = "shahh"

	fmt.Println(s1)

	// 值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["beijing"] = []int{10, 20, 30}
	fmt.Println(m1)

	//统计单词出现的次数 how do you do how=1 do =2 you=1
	str := "how do you do"
	ret := strings.Split(str, " ")
	fmt.Println(ret)
	m2 := make(map[string]int, 10)
	for _, v := range ret {
		_, ok := m2[v]
		if !ok {
			m2[v] = 1
		} else {
			m2[v] += m2[v]
		}
	}
	fmt.Println(m2)
}
