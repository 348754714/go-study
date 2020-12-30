package main

// 包的路径从GOPATH/src后面的路径开始写起
// 想被别的包调用的标识符首字母大写
// 单行导入和多行导入
// 导入包指定别名
// 导入包不想使用包内部的标识符，需要匿名导入
// 匿名导入的包只执行init()方法
// 导入包，但是不使用包中的方法，不能导入
import (
	"fmt" //官方包与第三方包习惯用空行分隔
	// . "fmt" .作为包别名时可以直接使用该包的方法，不推荐使用，

	// _ "github.com/348754714/go-study/day05/10calc" 匿名导入
	//"github.com/348754714/go-study/day05/10calc" // 文件夹的名字
	z "github.com/348754714/go-study/day05/10calc" //自定义包别名
)

// init执行时机
// 1 全局声名
// 2 init()
// main()
func init() {
	fmt.Println("自动执行...")
	fmt.Println(x, pi)
}

var x = 100

const pi = 3.14

// 导入包的init执行时机
// main import A
// A import B
// B import C
// C init()
// B init()
// A init()
// main init()

// 一般文件夹和包名一致
// go禁止循环导入包，A导入B，B再导入A
// 匿名导入包，只希望导入包，而不使用包中的数据，匿名导入的包也会编译到可执行文件中
// import _ "包路径" 只会执行该包中的init方法
// 导入包会立即执行init()方法
func main() {
	a := z.Add(1, 2) //包的名字
	fmt.Println(a)
}
