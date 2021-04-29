package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

// 暴力穷举 超出时间限制
//func minDistance1(word1 string, word2 string) int {
//	return helper(len(word1)-1, len(word2)-1, word1, word2)
//}
//
//func helper(i int, j int, word1 string, word2 string) int {
//	if i < 0 {
//		return j + 1
//	}
//
//	if j < 0 {
//		return i + 1
//	}
//
//	if word1[i] == word2[j] {
//		return helper(i-1, j-1, word1, word2) // 啥都不做
//	} else {
//		return min(min(
//			helper(i-1, j, word1, word2)+1,  // 删除
//			helper(i, j-1, word1, word2)+1), // 插入
//			helper(i-1, j-1, word1, word2)+1) // 替换
//	}
//}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// base case
	for i := 1; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < n+1; i++ {
		dp[0][i] = i
	}

	// 自底向上
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(
					dp[i-1][j]+1,  // 删除
					dp[i][j-1]+1), // 插入
					dp[i-1][j-1]+1) // 替换
			}
		}
	}

	return dp[m][n]
}
