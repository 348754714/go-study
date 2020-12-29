package main

import "fmt"

// 结构体模拟实现继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s can move!", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%s在旺旺\n", d.animal.name)
	fmt.Printf("ta有%d条腿\n", d.feet)
}

func main() {
	dog := dog{
		feet:   4,
		animal: animal{name: "dog"},
	}
	dog.wang()
	dog.move()
	//fmt.Println(dog.animal.name)
}
