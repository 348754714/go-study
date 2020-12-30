package main

import "fmt"

// 接口的实现
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

func (c chicken) move() {
	fmt.Println("鸡动")
}

func (c chicken) eat(food string) {
	fmt.Println("鸡吃", food)
}

func (c cat) move() {
	fmt.Println("猫步")
}

func (c cat) eat(food string) {
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
	fmt.Printf("%T\n", a1)
	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc // 结构体赋值给了接口
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", bc)
	cc := chicken{
		feet: 2,
	}
	move(a1)
	eat(a1, "鱼")
	move(cc)
	eat(cc, "饲料")

}
