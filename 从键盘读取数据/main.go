package main

import (
	"fmt"
	"math"
)

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

func mergeSort(a []int) {
	sort(a, 0, len(a)-1)
}

func sort(a []int, l, h int) {
	if l >= h {
		return
	}
	m := l + (h-l)/2
	sort(a, l, m)
	sort(a, m+1, h)
	merge(a, l, m, h)
}

func merge(a []int, l int, m int, h int) {
	b := make([]int, h-l+1)
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

	for i, k := l, 0; i <= h; i, k = i+1, k+1 {
		a[i] = b[k]
	}
}

func mergeSort1(a []int) {
	n := len(a)
	for sz := 1; sz < n; sz = 2 * sz {
		for l := 0; l < n-sz; l = l + 2*sz {
			merge(a, l, l+sz-1, int(math.Min(float64(n-1), float64(l+2*sz-1))))
		}
	}
}

func quickSort(a []int) {
	realQuickSort(a, 0, len(a)-1)
}

func realQuickSort(a []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	m := partition(a, lo, hi)
	realQuickSort(a, lo, m-1)
	realQuickSort(a, m+1, hi)
}

func partition(a []int, lo int, hi int) int {
	v := a[lo]
	i, j := lo+1, hi

	for {
		for a[i] <= v && i < hi {
			i++
		}

		for a[j] >= v && j > lo {
			j--
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	a[j], a[lo] = a[lo], a[j]
	return j
}

func quick3way(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	lt, gt := lo, hi
	i := lt + 1
	v := a[lo]

	for i <= gt {
		if a[i] < v {
			a[i], a[lt] = a[lt], a[i]
			lt++
			i++
		} else if a[i] > v {
			a[i], a[gt] = a[gt], a[i]
			gt--
		} else {
			i++
		}
	}
	quick3way(a, lo, lt-1)
	quick3way(a, gt+1, hi)
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	shellSort(a)
	fmt.Println(a)

	a = []int{5, 4, 3, 2, 1}
	mergeSort(a)
	fmt.Println(a)

	a = []int{5, 4, 3, 2, 1}
	mergeSort1(a)
	fmt.Println(a)

	a = []int{5, 4, 3, 2, 1}
	quickSort(a)
	fmt.Println(a)

	a = []int{5, 4, 3, 2, 1, 3, 2, 1}
	quick3way(a, 0, len(a)-1)
	fmt.Println(a)
}
