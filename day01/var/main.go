package main

import "fmt"

//  声名变量 推荐使用小驼峰
//var name string
//var age int
//var isOk bool

// 批量声名
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	//var heihei string // 变量声明必须使用
	fmt.Print(isOk)             // 普通打印
	fmt.Println() // 打印空行
	fmt.Printf("name:%s", name) // 占位符打印
	fmt.Println(age)            // 打印后添加换行

	// 声名同时赋值
	var s1 string = "hhh"
	fmt.Println(s1);
	//s1 := "10" // 同个作用域中不能声名同名变量

	// 类型推导声名
	var s2 = "h好搞笑"
	fmt.Println(s2)

	// 简短声名 只能再函数中使用
	s3 := "哈希"
	fmt.Println(s3)
	// 匿名变量 _ 接受不需要用到的函数返回值
}
