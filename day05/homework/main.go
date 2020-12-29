package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎光临学生管理系统")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.修改学生
		4.删除学生
		5.退出
		`)
}

var stuMgr studentManage

//var stuMgr = studentManage{
//	allStudent: make(map[int64]student, 100),
//}

func main() {
	stuMgr = studentManage{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		var input int
		fmt.Scanln(&input)
		fmt.Printf("您选择了%d选项\n", input)
		// 2.等待用户选择下一步
		// 3.用户选择后执行对应函数
		switch input {
		case 1:
			stuMgr.all()
		case 2:
			stuMgr.add()
		case 3:
			stuMgr.edit()
		case 4:
			stuMgr.del()
		case 5:
			os.Exit(1) // 退出
		default:
			fmt.Println("选择错误，请重新选择")
		}
	}
}
