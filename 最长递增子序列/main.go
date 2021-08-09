package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

// 使用动态规划
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxValue := 0
	for i := 0; i < len(nums); i++ {
		maxValue = max(maxValue, dp[i])
	}

	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
