package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age" db:"age" ini:"age"`
}

func main() {
	p1 := person{
		Name: "周关注",
		Age:  10,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("failed,err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	// 反序列化
	str := `{"name":"name1","age":18}`
	var p2 person
	err = json.Unmarshal([]byte(str), &p2) // 传指针是为了在json.Unmarshal内部修改p2的值
	if err != nil {
		fmt.Printf("Unmarshal,err:%v", err)
		return
	}
	fmt.Printf("%v\n", p2)
}
