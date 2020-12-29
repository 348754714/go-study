package main

import (
	"fmt"
	"os"
)

/*
函数版学生管理系统
查看，新增，删除学生
*/

var (
	allStudent map[uint64]*student
)

type student struct {
	id   uint64
	name string
	age  uint8
}

// 显示全部
func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s 年龄:%d\n", k, v.name, v.age)
	}
}

func newStudent(id uint64, name string, age uint8) *student {
	return &student{
		id:   id,
		name: name,
		age:  age,
	}
}

// 添加
func addStudent() {
	var (
		id   uint64
		name string
		age  uint8
	)
	fmt.Print("请输入学号")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄")
	fmt.Scanln(&age)
	newstu := newStudent(id, name, age)
	allStudent[id] = newstu
}

// 删除
func deleteStudent() {
	var (
		id uint64
	)
	fmt.Print("请输入学号删除学生：")
	fmt.Scanln(&id)
	delete(allStudent, id)
}
func main() {
	allStudent = make(map[uint64]*student, 100)
	// 1.打印菜单
	fmt.Println("欢迎光临学生管理系统")

	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出系统
		`)
	for {
		var input int
		fmt.Scanln(&input)
		fmt.Printf("您选择了%d选项\n", input)
		// 2.等待用户选择下一步
		// 3.用户选择后执行对应函数
		switch input {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出
		default:
			fmt.Println("选择错误，请重新选择")
		}
	}
}
