package main

//221. 最大正方形
// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
// 暴力求解，时间复杂度很高
// 动态规划。
// dp[i][j]表示以i,j为右下角的最大的边,说明右下角都为1
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])
	maxSide := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// base case
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxSide = 1
		}
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxSide = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
