package main

import "fmt"

// 结构体实现构造函数
type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

// 构造函数返回的是结构体还是结构体指针
// 当架构体较大的时候尽量使用结构体指针减少内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name,
	}
}

func main() {
	p1 := newPerson("name1", 20)
	p2 := newPerson("name2", 10)
	fmt.Println(*p1, *p2)

	d := newDog("dahuang")
	fmt.Println(d)
}
