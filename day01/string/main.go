package main

import (
	"fmt"
	"strings"
)

// 字符串
//字符串使用双引号包裹
//s := "dsahfd"
//字符用单引号包裹
//c := 'f'
//字符串转义
//\r 回车返回首行
//\n 换行直接跳到下一行的同列位置
//\t tab制表符
//\' 单引号
//\" 双引号
//\\ 反斜杠
// 常用操作
//len(srt) 求长度
//+或fmt.Sprintf 拼接
//strings.Split 分割
//strings.contains 是否包含
//strings.HasPrefix,strings.HasSuffix 是否包含前缀，后缀
//strings.Index(),strings.LastIndex() 子串第一次和最后一次出现的位置
//strings.Join(a[] string, sep string) join操作
func main() {
	path := "D:\\work\\project\\go\\src\\github.com\\348754714"
	fmt.Printf("path = \"%s\"", path)

	s := "I'm OK"
	fmt.Println(s)
	// 多行字符串``包裹，原样输出
	s2 := `
世情薄
			人情恶
		雨送黄昏花易落

`
	fmt.Println(s2)
	s3 := `D:\work\project\go\src\github.com\348754714`
	fmt.Printf("path = \"%s\"\n", s3)

	fmt.Println(len(s3))
	fmt.Println(s + s3)
	ss1 := fmt.Sprintf("%s%s", s, s3)
	fmt.Println(ss1)

	// 分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(s3, "work"))
	fmt.Println(strings.Contains(s3, "works"))

	// 前后缀
	fmt.Println(strings.HasPrefix(s3, "D:"))
	fmt.Println(strings.HasPrefix(s3, "D1:"))
	fmt.Println(strings.HasSuffix(s3, "D1:"))
	fmt.Println(strings.HasSuffix(s3, "348754714"))

	// 位置
	fmt.Println(strings.Index(s3, "348754714"))
	fmt.Println(strings.LastIndex(s3, "o"))
	fmt.Println(strings.Index(s3, "dai")) // -1
	fmt.Println(strings.Index(s3, "D"))   // 0

	// join
	fmt.Println(strings.Join(ret, "+"))
}
