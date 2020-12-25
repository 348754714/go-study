package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map
// map 是引用类型，使用前必须申请内存空间
func main() {
	var m1 map[string]int         // 没有初始化没有分配内存空间
	m1 = make(map[string]int, 10) // 要估算好map空间，避免在程序运行期间再动态扩容
	m1["age"] = 18
	m1["jiwum"] = 35

	fmt.Println(m1)
	fmt.Println(m1["age"])
	fmt.Println(m1["jiwum1"]) // 如果不存在返回KEY对应的0值
	// 约定成俗用ok接收bool
	value, ok := m1["jiwum1"]
	if !ok {
		fmt.Println("no key")
	} else {
		fmt.Println(value)
	}

	// 遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 遍历KEY
	for k := range m1 {
		fmt.Println(k)
	}
	// 遍历VALUE
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除
	delete(m1, "jiwum")
	fmt.Println(m1)
	delete(m1, "fdhjsaf") // 删除不存在的

	// 看内置文档命令行
	//go doc builtin.delete
	// go 中文文档
	// https://studygolang.com/pkgdoc

	// 按照指定的顺序遍历map
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := rand.Intn(100)          // 0-99随机整数
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)
	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
