package main

// 121. 买卖股票的最佳时机
// dp[i][1]第i天持有股票的获利
// dp[i][0]第i天没有股票的获利
func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < size; i++ {
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = maxInt(dp[i-1][1], -prices[i])
	}
	return dp[size-1][0]
}
