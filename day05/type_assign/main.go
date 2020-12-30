package main

import "fmt"

//类型断言
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println(str)
	}
}

func switch_assign(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) { // 拿到接口中值的类型
	case string:
		fmt.Printf("%T\n", t)
	case int:
		fmt.Printf("%T\n", t)
	case int64:
		fmt.Printf("%T\n", t)
	case bool:
		fmt.Printf("%T\n", t)
	}
}

func main() {
	//assign("fdsjkfsd")
	switch_assign(100)
	switch_assign("fdhsjfs")
	switch_assign(int64(1232312321))
	switch_assign(false)
}
