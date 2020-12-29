package main

import "fmt"

// 结构体学生管理系统

type student struct {
	id   int64
	name string
}

type studentManage struct {
	allStudent map[int64]student
	student
}

func (s *studentManage) add() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名")
	fmt.Scanln(&name)
	newstu := newStudent(id, name)
	s.allStudent[id] = *newstu
}

func (s *studentManage) del() {
	var (
		id int64
	)
	fmt.Print("请输入学号删除学生：")
	fmt.Scanln(&id)
	_, ok := s.allStudent[id]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudent, id)
}

func (s *studentManage) edit() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号修改学生：")
	fmt.Scanln(&id)
	info, ok := s.allStudent[id]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("学号：%d 姓名：%s", info.id, info.name)
	fmt.Print("请修改学生姓名：")
	fmt.Scanln(&name)
	info.name = name
	s.allStudent[id] = info
	//delete(s.allStudent, id)
}

func (s *studentManage) all() {
	allStudent := s.allStudent
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
