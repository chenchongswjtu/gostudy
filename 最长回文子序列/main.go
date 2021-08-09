package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}

// dp[i][j]表示s[i:j+1]的最长回文子序列的长度
// dp[i][j]表示s[i:i+1]一个字符的最长回文子序列为1
// 最后返回dp[0][n-1]
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// 倒着计算
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
