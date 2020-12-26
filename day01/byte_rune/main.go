package main

import (
	"fmt"
	"unicode"
)

// byte（uint8）存储 和 rune（int32） 类型 代表一个UTF8存储
// GO语言为了处理非ASCII类型支付，定义了rune类型
func main() {
	s := "Hello 戴老大 맏이를거느리다"
	n := len(s)
	fmt.Println(n)

	//for i := 0; i < n; i++ {
	//	fmt.Println(s[i]) // 无法处理中文
	//	//fmt.Printf("%c\n", s[i]) // 无法处理中文
	//}
	// 判断汉字数量
	var count int
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
		fmt.Printf("%c\n", c) // 处理中文
	}

	fmt.Println(count)

	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 将字符串强制转换成rune切片
	s3[0] = '红'
	fmt.Println(string(s3))
	fmt.Println(s3)
	fmt.Printf("%c%c%c\n", s3[0], s3[1], s3[2])

	c1 := "H"
	c2 := 'H'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	n1 := 100
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

}
