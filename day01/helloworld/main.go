package main

import "fmt"

//Mac 下编译 Linux 和 Windows平台 64位 可执行程序：
//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
//CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
//
//Linux 下编译 Mac 和 Windows 平台64位可执行程序：
//CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
//CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
//
//Windows下编译Mac平台64位可执行程序：
//SET CGO_ENABLED=0
//SET GOOS=darwin
//SET GOARCH=amd64
//go build

//Windows下编译linux平台64位可执行程序：
//SET CGO_ENABLED=0
//set GOARCH=amd64
//set GOOS=linux
//go build

// 函数外面只能放置标识符（变量，常量，函数，类型）的声名
//go build 当前目录可执行文件
//go build -o "a.exe" 编译并指定文件名称
//go run main.go 脚本执行
//go install 先编译后拷贝
func main() {
	fmt.Println("Hellor world!")
}
