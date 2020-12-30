package main

import "fmt"

// 使用值接收者和指针接收者的区别
// 使用值接收者实现方法，值和指针都能传
// 使用指针接收者实现方法，只能传指针
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c *cat) move() {
	fmt.Println("猫步")
}

func (c *cat) eat(food string) {
	fmt.Println("猫吃", food)
}

func move(a animal) {
	a.move()
}

func eat(a animal, food string) {
	a.eat(food)
}

func main() {
	var a1 animal
	c1 := cat{"tom", 4}
	c2 := &cat{"假老练", 4}

	a1 = &c1 // 实现animal接口的是cat的指针类型
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
