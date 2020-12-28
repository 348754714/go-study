package main

import (
	"fmt"
)

// GO中如果标识符首字母大写，就表示外部包可见（暴露的，公有的）
// Dog 是要给外部可见的结构体
type Dog struct {
	name string
}

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name,
	}
}

type person struct {
	name string
	age  int
}

// 值接收，传拷贝
func (p person) guonian() {
	p.age++
}

// 指针接收者：传地址
// 什么时候使用
// 1 需要修改接收者中的值
// 2 接收者是拷贝代价比较大的大对象
// 3 保持一致性，如果由某个方法使用了指针接收者，那么其他方法也应该使用指针接收者
func (p *person) zhenguonian() {
	p.age++
}

func newPerson(name string, age int) person {
	return person{
		name,
		age,
	}
}

// 方法作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示 d dog
func (d dog) wang(s string) {
	fmt.Printf("%s:%s", d.name, s)
}

// 方法
func main() {
	d1 := newDog("huangpi")
	// 类似于class的成员函数
	d1.wang("旺旺")

	p1 := newPerson("name1", 18)
	fmt.Println(p1.age) // 18
	p1.guonian()
	fmt.Println(p1.age) // 18
	p1.zhenguonian()
	fmt.Println(p1.age) // 19
}
