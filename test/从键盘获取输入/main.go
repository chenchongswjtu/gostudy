package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 测试怎么从键盘输入数据
func main() {
	var l int
	// fmt.Println("input len")
	_, err := fmt.Scanln(&l)
	if err != nil {
		fmt.Println(err)
		return
	}

	var n int
	// fmt.Println("input num")
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(l, n)

	// 创建二维数组
	var arrs = make([][]int, n)
	for i := 0; i < n; i++ {
		var str string
		// 从键盘输入字符串（每个数组以，分开）
		_, err = fmt.Scanln(&str)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 以，将字符串切分为字符串数组
		ss := strings.Split(str, ",")
		var arr = make([]int, len(ss))
		// 将字符串数组转换为int
		for j, s := range ss {
			k, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				return
			}
			arr[j] = k
		}
		arrs[i] = arr
	}

	fmt.Println(arrs)

	// todo
	//var a, b int
	//_, err = fmt.Scanf("%d,%d", &a, &b)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(a, b)
}
