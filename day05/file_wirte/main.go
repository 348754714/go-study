package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件写入

func file_wirte() {
	fileObj, err := os.OpenFile("./1.file", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed:%s", err)
		return
	}
	defer fileObj.Close()

	// write
	fileObj.Write([]byte("fhdjsafdsa\n"))
	fileObj.WriteString("ffdhjsfdjsafhjdsaf\n")
}

func file_wirte1() {
	fileObj, err := os.OpenFile("./1.file", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed:%s", err)
		return
	}
	defer fileObj.Close()

	// write
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("fdjsfdskafds\n")
	wr.Flush() // 缓冲区写入文件
}

func file_wirte2() {
	str := "fjhhhhhhhhhhh"
	err := ioutil.WriteFile("./1.file", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed:%v\n", err)
		return
	}
}
func main() {
	file_wirte()
	file_wirte1()
	file_wirte2()
	return
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed:%s", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()

	// 读文件
	//var tmp = make([]byte, 128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read file failed:%v\n", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}
