package main

import (
	"fmt"
	"math"
)

func mergeSort(a []int) {
	b := make([]int, len(a))
	sort(a, 0, len(a)-1, b)
}

func sort(a []int, l int, h int, b []int) {
	if h <= l {
		return
	}
	m := l + (h-l)/2
	sort(a, l, m, b)
	sort(a, m+1, h, b)
	merge(a, l, m, h, b)
}

func merge(a []int, l int, m int, h int, b []int) {
	i, j, k := l, m+1, 0

	for i <= m && j <= h {
		if a[i] <= a[j] {
			b[k] = a[i]
			i++
			k++
		} else {
			b[k] = a[j]
			j++
			k++
		}
	}

	for i <= m {
		b[k] = a[i]
		i++
		k++
	}

	for j <= h {
		b[k] = a[j]
		j++
		k++
	}

	k = 0
	for i = l; i <= h; {
		a[i] = b[k]
		i++
		k++
	}
}

func mergeSort1(a []int) {
	n := len(a)
	b := make([]int, len(a))
	for sz := 1; sz < len(a); sz = sz + sz {
		for l := 0; l < n-sz; l = l + sz + sz {
			merge(a, l, l+sz-1, int(math.Min(float64(l+sz+sz-1), float64(n-1))), b)
		}
	}
}

func main() {
	var a = []int{5, 4, 3, 2, 1, 0, 9, 8}
	mergeSort1(a)
	fmt.Println(a)
	a = []int{5, 4, 3, 2, 1, 0, 9, 8}
	mergeSort(a)
	fmt.Println(a)
}
