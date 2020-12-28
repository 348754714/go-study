package main

import "fmt"

// 结构体占用一块连续的内存空间
type x struct {
	a int8
	b int8
	c int8
	d int8
	e string
	f int8
}

func main() {
	m := x{
		a: int8(1),
		b: int8(2),
		c: int8(3),
		d: int8(4),
		e: "fdsafhdjsaf",
		f: int8(5),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
	fmt.Printf("%p\n", &(m.d))
	fmt.Printf("%p\n", &(m.e))
	fmt.Printf("%p\n", &(m.f))
}
