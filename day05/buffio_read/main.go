package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// buffio读文件
func readFromFilebuBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed:%s", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()

	// 创建一个用来从文件中读的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n') // \n 行读取
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed:%v\n", err)
		}
		fmt.Print(line)
	}

}

// ioutil读文件
func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

// 打开文件

func main() {
	//readFromFilebuBufio()
	readFromFileByIoutil()
}
