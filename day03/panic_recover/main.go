package main

import "fmt"

// panic 和 recover
func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		// recover必须搭配defer使用
		// defer必须在可能引发panic之前定义
		err := recover() // 尝试恢复错误然后继续执行，少用
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现严重级别的错误")
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
