package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var length int
	var num int
	_, err := fmt.Scanln(&length)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Scanln(&num)
	if err != nil {
		fmt.Println(err)
		return
	}

	var inputs = make([][]int, num)
	var sum int
	for i := 0; i < num; i++ {
		var str string
		_, err = fmt.Scanln(&str)
		if err != nil {
			fmt.Println(err)
			return
		}
		ss := strings.Split(str, ",")
		tmp := make([]int, len(ss))
		for j, s := range ss {
			v, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				return
			}
			tmp[j] = v
		}
		inputs[i] = tmp
		sum += len(tmp)
	}

	// 逻辑
	var ret []int

	for {
		if len(ret) == sum {
			break
		}

		for i := 0; i < len(inputs); i++ {
			if len(inputs[i]) == 0 {
				continue
			}
			size := minInt(length, len(inputs[i]))
			tmp := inputs[i][:size]
			ret = append(ret, tmp...)
			inputs[i] = inputs[i][size:]
		}
	}

	//fmt.Println(ret)
	var res = make([]string, len(ret))
	for i, v := range ret {
		res[i] = strconv.Itoa(v)
	}

	fmt.Println(strings.Join(res, ","))

}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
