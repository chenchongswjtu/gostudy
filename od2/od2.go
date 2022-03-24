package main

import "fmt"

func main() {
	var per int
	fmt.Scanln(&per)
	if per < 1 || per > 10000 {
		return
	}

	var n int
	fmt.Scanln(&n)
	if n < 1 || n > 10000 {
		return
	}

	var task = make([]int, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		if t < 1 || t > 10000 {
			return
		}
		task[i] = t
	}

	//fmt.Println(task)

	var cost int
	var left int
	for i := 0; i < n; i++ {
		if task[i]+left <= per {
			cost++
			left = 0
		} else {
			cost++
			left = (task[i] + left) - per
		}
	}

	if left%per == 0 {
		cost += left / per
	} else {
		cost += left/per + 1
	}

	fmt.Println(cost)
}
