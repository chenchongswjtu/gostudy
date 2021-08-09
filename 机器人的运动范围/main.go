package main

import "fmt"

func movingCount(m int, n int, k int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	return dfs(m, n, 0, 0, k, dp)
}

func dfs(m, n, i, j, k int, dp [][]int) int {
	if i < 0 || j < 0 || i >= m || j >= n || dp[i][j] == 1 || (sumPos(i)+sumPos(j)) > k {
		return 0
	}

	dp[i][j] = 1

	sum := 1
	sum += dfs(m, n, i, j+1, k, dp)
	sum += dfs(m, n, i, j-1, k, dp)
	sum += dfs(m, n, i+1, j, k, dp)
	sum += dfs(m, n, i-1, j, k, dp)
	return sum
}

// 求所有位之和
func sumPos(n int) int {
	var sum int

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}

func main() {
	fmt.Println(movingCount(10, 10, 17))
}
