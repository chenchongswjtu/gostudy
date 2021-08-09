package main

import (
	"fmt"
)

// 希尔排序，在插入排序的基础上进行优化
func shellSort(a []int) {
	n := len(a)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h = h / 3
	}
}

func main() {
	var a = []int{5, 4, 3, 2, 1, 0, 9, 8}
	shellSort(a)
	fmt.Println(a)
}
