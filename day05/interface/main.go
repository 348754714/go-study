package main

import "fmt"

// 引出接口的实例
// 接口是一种特殊类型，它规定了变量有哪些方法
// 不关心变量是什么类型，只关心能调用变量的方法
type cat struct {
}

type dog struct {
}

type person struct {
}

// 只要实现了speak()方法的变量都是speaker类型
type speaker interface {
	speak() // 方法签名，可以一个也可以多个
}

func (c cat) speak() {
	fmt.Println("喵~")
}

func (d dog) speak() {
	fmt.Println("汪~")
}

func (p person) speak() {
	fmt.Println("哎呦~")
}

func da(x speaker) {
	// 接收一个参数，传进来什么就打什么
	x.speak() // 挨打就要叫
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	c1.speak()
	d1.speak()
	p1.speak()
}
