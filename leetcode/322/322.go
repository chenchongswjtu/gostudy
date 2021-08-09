package main

import "fmt"

// 动态规划，换零钱
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 0; j < amount; j++ {
			if dp[j] != 0 && j+coins[i] <= amount {
				dp[j+coins[i]] += dp[j]
			}

		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
