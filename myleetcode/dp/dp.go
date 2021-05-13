package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numDecodings("226"))
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

// 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// 64. 最小路径和
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

// 97. 交错字符串
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}

	f := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		f[i] = make([]bool, n2+1)
	}
	f[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			k := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[k])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[k])
			}
		}
	}

	return f[n1][n2]
}

// 120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	m, n := len(triangle), len(triangle[len(triangle)-1])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[n-1][i] = triangle[n-1][i]
	}

	for i := m - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

// 121. 买卖股票的最佳时机
// 只能买卖一次
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	min := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}

// 122. 买卖股票的最佳时机II
// 买卖次数不限
// 定义状态 dp[i][0] 表示第 ii 天交易完后手里没有股票的最大利润
// dp[i][1] 表示第 ii 天交易完后手里持有一支股票的最大利润（ii 从 00 开始）
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	n := len(prices)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// 123. 买卖股票的最佳时机 III
func maxProfit3(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

// 124. 买卖股票的最佳时机 IV
// 因此对 dp[i][k] 的定义需要分成两项：

// dp[i][k][0] 表示在第 i 天结束时，最多进行 k 次交易且在进行操作后持有 0 份股票的情况下可以获得的最大收益；
// dp[i][k][1] 表示在第 i 天结束时，最多进行 k 次交易且在进行操作后持有 1 份股票的情况下可以获得的最大收益。
func maxProfit4(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(prices)
	if k >= n/2 { //如果股票价格数组的长度为 n，则有收益的交易的数量最多为 n / 2（整数除法）。因此 k 的临界值是 n / 2。如果给定的 k 不小于临界值，即 k >= n / 2，则可以将 k 扩展为正无穷，此时问题等价于情况二。
		return maxProfit2(prices)
	}

	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}

	for i := 1; i <= k; i++ {
		dp[0][i][1] = -prices[0]
		dp[0][i][0] = 0
	}

	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
		}
	}

	return dp[n-1][k][0]
}

// 131. 分割回文串
func partition(s string) [][]string {
	var ans [][]string
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	ss := make([]string, 0)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), ss...))
			return
		}

		for j := i; j < n; j++ {
			if dp[i][j] {
				ss = append(ss, s[i:j+1])
				dfs(j + 1)
				ss = ss[:len(ss)-1]
			}
		}
	}

	dfs(0)
	return ans
}

// 91. 解码方法(回溯解法超时)
func numDecodings1(s string) int {
	var count int

	isValid := func(s string) bool {
		if s[0] == '0' {
			return false
		}

		var res int
		for i := range s {
			res = res*10 + int(s[i]-'0')
		}
		if res >= 1 && res <= 26 {
			return true
		}
		return false
	}

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(s) {
			count++
			return
		}
		for i := index; i < len(s); i++ {
			t := s[index : i+1]
			if isValid(t) {
				dfs(i + 1)
			} else {
				break
			}
		}
	}

	dfs(0)
	return count
}

// 91. 解码方法(动态规划解法)
func numDecodings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

// 152. 乘积最大子数组
func maxProduct(nums []int) int {
	n := len(nums)
	maxf := make([]int, n)
	minf := make([]int, n)
	copy(maxf, nums)
	copy(minf, nums)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		maxf[i] = max(maxf[i-1]*nums[i], max(nums[i], minf[i-1]*nums[i]))
		minf[i] = min(minf[i-1]*nums[i], min(nums[i], maxf[i-1]*nums[i]))
	}
	var ans = maxf[0]
	for _, m := range maxf {
		if m > ans {
			ans = m
		}
	}
	return ans
}

// 213. 打家劫舍 II
func rob2(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	var _rob func(nums []int) int
	_rob = func(nums []int) int {
		first, second := nums[0], max(nums[0], nums[1])
		for _, v := range nums[2:] {
			first, second = second, max(first+v, second)
		}
		return second
	}

	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

// 221. 最大正方形(动态规划)
// dp(i,j) 表示以 (i,j) 为右下角，且只包含 1 的正方形的边长最大值
// dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

// 264. 丑数 II
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
