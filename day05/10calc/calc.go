package calc

import "fmt"

func init() {
	fmt.Println("import 我的时候自动执行......")
}

// 首字母小写表示私有
func Add(x, y int) int {
	return x + y
}
