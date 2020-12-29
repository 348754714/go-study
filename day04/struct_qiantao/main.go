package main

import "fmt"

// 结构体嵌套

type school struct {
	name    string
	address string
}
type student struct {
	id   uint64
	name string
	age  uint8
}

type student_school struct {
	gender  string
	student student
	school  school
}

func main() {
	student_school := student_school{
		gender: "男",
		student: student{
			id:   100,
			name: "name1",
			age:  10,
		},
		school: school{
			name:    "school1",
			address: "address1",
		},
	}
	fmt.Println(student_school)
	fmt.Println(student_school.school)
	fmt.Println(student_school.school.name)
	fmt.Println(student_school.student.age)
}
