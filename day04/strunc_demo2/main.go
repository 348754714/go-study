package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "女" // 修改的是副本person的gender
}

func f2(x *person) {
	//(*x).gender = "女" // 根据指针找到变量，修改原来的变量
	x.gender = "女" // 语法糖，自动根据指针找到对应的变量
}

func main() {
	var p1 person
	p1.name = "TOM"
	p1.gender = "male"
	f(p1)
	fmt.Println(p1.gender) //male

	f2(&p1)
	fmt.Println(p1)

	// 创建结构体类型指针
	var p2 = new(person)
	//(*p2).name = "PTOM"
	(*p2).name = "PTOM"
	p2.gender = "male" // 语法糖
	f2(p2)
	fmt.Println(*p2)
	fmt.Printf("%p\n", p2)

	// 结构体指针2,k-v赋值，常用方式
	var p3 = person{
		name:   "TOM3",
		gender: "男",
	}
	fmt.Println(p3)

	p4 := person{ // 按照顺序赋值
		"TOM4",
		"男",
	}
	fmt.Println(p4)
}
