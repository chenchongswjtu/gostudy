package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDrop(5, 14))
}

func superEggDrop(k int, n int) int {
	memo := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		memo[i] = make([]int, n+1)
	}

	for i := 1; i < k+1; i++ {
		for j := 1; j < n+1; j++ {
			memo[i][j] = -1
		}
	}

	return dp(memo, k, n)
}

func dp(memo [][]int, k int, n int) int {
	if k == 1 {
		return n
	}

	if n == 1 {
		return 1
	}

	if memo[k][n] != -1 {
		return memo[k][n]
	}

	res := math.MaxInt64
	for i := 1; i <= n; i++ {
		res = min(res, max(dp(memo, k-1, i-1), dp(memo, k, n-i))+1)
	}

	memo[k][n] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
