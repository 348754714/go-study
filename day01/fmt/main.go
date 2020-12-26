package main

import "fmt"

// fmt 占位符
func main() {
	var n = 100
	fmt.Printf("%T\n", n) // 打印类型
	fmt.Printf("%v\n", n) // 打印值
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%d\n", n) // 10进制
	fmt.Printf("%o\n", n) // 8 进制
	fmt.Printf("%x\n", n) // 16进制
	fmt.Printf("%s\n", n) // 打印字符串，整型打印字符串非想要的结果
	var s = "HELLO"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)  // HELLO
	fmt.Printf("%#v\n", s) // "HELLO"

	fmt.Printf("%q\n", 65)
	// 浮点数和复数
	fmt.Printf("%b\n", 3.1415926)
	// 字符串
	fmt.Printf("%q\n", "fdhjsaf")

	f := 12.34
	fmt.Printf("%f\n", f)
	fmt.Printf("%9f\n", f)    // 宽度9，默认精度
	fmt.Printf("%.2f\n", f)   // 默认宽度，精度2
	fmt.Printf("%09.2f\n", f) //左填充0宽度9，精度2
	fmt.Printf("%9.f\n", f)   // 宽度9，精度0
	// 获取用户输入， 读取由空白字符分割的值保存并传递给本函数的参数，换行符视为空白
	//var ss string
	//fmt.Scan(&ss)
	//fmt.Println(ss)

	var (
		name  string
		age   int
		class string
	)
	//fmt.Scanf("%s %d %s\n", &name, &age, &class)
	//fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
