package main

import "fmt"

// 接口示例2
// 不管是什么牌子的车都能跑

type car interface {
	run()
	sell(int)
}

type falali struct {
	brand string
}

type maibahe struct {
	brand string
}

type baoshijie struct {
	brand string
}

func (b baoshijie) run() {
	fmt.Printf("%s速度700\n", b.brand)
}

func (b baoshijie) sell(price int) {
	fmt.Printf("%s售价%d\n", b.brand, price)
}

func (f falali) run() {
	fmt.Printf("%s速度500\n", f.brand)
}

func (f falali) sell(price int) {
	fmt.Printf("%s售价%d\n", f.brand, price)
}

// 必须实现接口中定义的全部方法，才能调用car接口
func (m maibahe) run() {
	fmt.Printf("%s速度700\n", m.brand)
}

func drive(car1 car) {
	car1.run()
}

func seller(car1 car, price int) {
	car1.sell(price)
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	//var m1 = maibahe{
	//	brand: "迈巴赫",
	//}
	drive(f1)
	seller(f1, 1000000)
	drive(b1)
	seller(b1, 20000000)

}
