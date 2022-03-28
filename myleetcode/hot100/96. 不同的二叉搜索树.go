package main

//96. 不同的二叉搜索树
//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
//示例 1：
//
//
//输入：n = 3
//输出：5
//示例 2：
//
//输入：n = 1
//输出：1
//
//
//提示：
//
//1 <= n <= 19

// 动态规划
func numTrees(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += dp[j] * dp[i-1-j]
		}
		dp[i] = sum
	}

	return dp[n]
}
