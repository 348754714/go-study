package main

import (
	"fmt"
	"strings"
)

/* 50枚金币，分配给：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth
分配规则
a:包含1个e或E分1枚
b:包含1个i或I分2
c:包含1个o或O分3
d:包含1个u或U分4个
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distirbution = make(map[string]int, len(users))
)

func main() {
	remain := dispatchCoin2()
	fmt.Println(remain)
}

func dispatchCoin() int { // 原始写法
	dispatchByCharac := func(name string) int { // 根据字符返回
		count := 0
		count += strings.Count(name, "e")
		count += strings.Count(name, "E")
		count += strings.Count(name, "i") * 2
		count += strings.Count(name, "I") * 2
		count += strings.Count(name, "o") * 3
		count += strings.Count(name, "O") * 3
		count += strings.Count(name, "u") * 4
		count += strings.Count(name, "U") * 4
		return count
	}

	for _, v := range users {
		coin := dispatchByCharac(v)
		distirbution[v] = coin
		coins -= coin
	}

	fmt.Println(distirbution)
	return coins
}

func dispatchCoin1() int { // 写法1
	dispatchByCharac := func(name string, c string, reward int) int { // 根据字符返回
		count := 0
		lower := strings.ToLower(c)
		upper := strings.ToUpper(c)
		count += strings.Count(name, lower) * reward
		count += strings.Count(name, upper) * reward
		return count
	}

	for _, v := range users {
		coin := dispatchByCharac(v, "e", 1)
		coin += dispatchByCharac(v, "i", 2)
		coin += dispatchByCharac(v, "o", 3)
		coin += dispatchByCharac(v, "u", 4)
		distirbution[v] = coin
		coins -= coin
	}

	fmt.Println(distirbution)
	return coins
}

func dispatchCoin2() int { // 写法2
	c_reward := make(map[string]int, 26)
	c_reward["e"] = 1
	c_reward["i"] = 2
	c_reward["o"] = 3
	c_reward["u"] = 4
	dispatchByCharac := func(name string, c_reward map[string]int) int { // 根据字符返回
		count := 0
		for k, v := range c_reward {
			lower := strings.ToLower(k)
			upper := strings.ToUpper(k)
			count += strings.Count(name, lower) * v
			count += strings.Count(name, upper) * v
		}
		return count
	}

	for _, v := range users {
		coin := dispatchByCharac(v, c_reward)
		distirbution[v] = coin
		coins -= coin
	}

	fmt.Println(distirbution)
	return coins
}
