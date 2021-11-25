package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(LIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))

	fmt.Println(maxEnvelopes([][]int{{1, 8}, {5, 2}, {5, 4}, {6, 4}, {6, 7}, {2, 3}}))
}

// 最长递增子序列
func LengthOfLIS(nums []int) int {
	// dp[i] 表示以nums[i]为结尾的LIS长度
	dp := make([]int, len(nums))

	// 初始化长度为1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ret := 0
	for i := 0; i < len(dp); i++ {
		ret = max(ret, dp[i])
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 信封嵌套
// 根据w升序，w相同根据h降序
// 然后再算h的最长递增子序列的长度
// {{w,h},{w,h},......}
func maxEnvelopes(envs [][]int) int {
	sort.Slice(envs, func(i, j int) bool {
		if envs[i][0] == envs[j][0] {
			return envs[i][1] > envs[j][1]
		}

		return envs[i][0] < envs[j][0]
	})

	fmt.Println(envs)

	hs := make([]int, len(envs))
	for i := 0; i < len(envs); i++ {
		hs[i] = envs[i][1]
	}

	return LengthOfLIS(hs)
}

// 选择递增序列的长度和递增序列（不唯一）
func LIS(nums []int) (int, []int) {
	// dp[i] 表示以nums[i]为结尾的LIS长度
	dp := make([]int, len(nums))
	// preIndex[i] 表示以nums[i]结尾的递增序列的上一个节点索引,初始化为自己
	preIndex := make([]int, len(nums))

	// 初始化长度为1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		preIndex[i] = i
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					preIndex[i] = j // 更新上一个节点索引
				}
			}
		}
	}

	lastIndex := 0 // 递增序列的最后一个索引
	length := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > length {
			length = dp[i]
			lastIndex = i
		}
	}

	l := length
	res := make([]int, l)
	i := lastIndex
	for {
		l--
		res[l] = nums[i]
		if preIndex[i] == i { // 节点的上一个索引为自己，则介绍，找到第一个节点
			break
		}
		i = preIndex[i]
	}

	return length, res
}
