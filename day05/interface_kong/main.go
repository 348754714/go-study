package main

import "fmt"

// 空接口，没有必要起名字，通常定义成下面的格式
// 任意类型都实现了空接口，可以被任意类型调用
//interface {} 空接口类型

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {
	// 应用1，实现不固定类型的的数据格式
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "zhoulin"
	m1["age"] = 19
	m1["merried"] = true
	m1["hobby"] = [...]string{"chang", "tiao", "rap"}
	fmt.Println(m1)

	// 用用2，作为函数的参数（不能确定参数类型）
	show(false)
	show(nil)
	show(m1)

}
