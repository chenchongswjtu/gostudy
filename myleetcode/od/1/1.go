package main

import "fmt"

func main() {
	var n int
	var m int
	fmt.Scanln(&n)
	fmt.Scanln(&m)

	var count int
	for i := n; i < m; i++ {
		for j := n + 1; j < m; j++ {
			for k := n + 2; k < m; k++ {
				if i < j && j < k && k*k == i*i+j*j && huzhi(i, j) == 1 && huzhi(i, k) == 1 && huzhi(j, k) == 1 {
					fmt.Printf("%d %d %d\n", i, j, k)
					count++
				}
			}
		}
	}

	if count == 0 {
		fmt.Println("Na")
	}
}

// 互为质数的算法
func huzhi(a, b int) int {
	if a < b {
		a, b = b, a
	}

	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
