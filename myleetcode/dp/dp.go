package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(7, 3))
}

// 5. 最长回文子串
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	var maxLen = 1
	var ans = s[:1]
	// dp 二维数组表示s[i:j+1]是不是回文子串
	var dp = make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	update := func(i, j int) {
		if !dp[i][j] {
			return
		}
		if j-i+1 <= maxLen {
			return
		}
		maxLen = j - i + 1
		ans = s[i : j+1]
	}

	// 初始化s[i:i+1]一个字符肯定是回文子串
	for i := range dp {
		dp[i][i] = true
		j := i + 1
		if j < n && s[i] == s[j] {
			dp[i][j] = true
			update(i, j)
		}
	}

	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			update(i, j)
		}
	}
	return ans
}

func longestPalindrome1(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	var maxLen = 1
	var begin = 0
	// dp 二维数组表示s[i:j+1]是不是回文子串
	var dp = make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// 初始化s[i:i+1]一个字符肯定是回文子串
	for i := range dp {
		dp[i][i] = true
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

//53. 最大子序和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//53. 最大子序和
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

// 62. 不同路径
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
