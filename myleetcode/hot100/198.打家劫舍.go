package main

// 198.打家劫舍
// 动态规划
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return maxInt(nums[0], nums[1])
	}

	dp := make([]int, n)

	dp[0] = nums[0]
	dp[1] = maxInt(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = maxInt(dp[i-1], dp[i-2]+nums[i]) // 选择不打劫i最大为dp[i-1],打劫最大为dp[i-2]+nums[i]
	}

	return dp[n-1]
}
