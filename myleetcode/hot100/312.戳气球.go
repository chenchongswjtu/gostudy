package main

func maxCoins(nums []int) int {
	n := len(nums)
	vals := make([]int, n+2)
	vals[0] = 1 // 前后加1
	vals[n+1] = 1
	for i := 0; i < len(nums); i++ {
		vals[i+1] = nums[i]
	}

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	// i,k,j
	// i 0..n-1
	// j 2..n+1
	// k i+1,j
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				sum := vals[i] * vals[k] * vals[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = maxInt(dp[i][j], sum)
			}
		}
	}

	return dp[0][n+1]
}
