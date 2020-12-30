package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套
type animal interface {
	animal_move
	animal_eat
}
type animal_move interface {
	move()
}

type animal_eat interface {
	eat(string)
}

type cat struct {
	name string
	feet int
}

// 实现了animal_move接口
func (c *cat) move() {
	fmt.Println("猫步")
}

// 实现了animal_eat接口
func (c *cat) eat(food string) {
	fmt.Println("猫吃", food)
}

func main() {

}
