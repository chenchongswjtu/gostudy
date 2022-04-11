package main

//41. 缺失的第一个正数
//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
//
//示例 1：
//
//输入：nums = [1,2,0]
//输出：3
//示例 2：
//
//输入：nums = [3,4,-1,1]
//输出：2
//示例 3：
//
//输入：nums = [7,8,9,11,12]
//输出：1

// [3,4,-1,1]
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1 // 将小于等于0的数改为n+1,正常情况下不会出现的数，只会出现1到n，之后的数字都是大于0的
		}
	}

	for i := 0; i < n; i++ {
		index := abs(nums[i])
		if index <= n { // 将对应<=n的数对于的值变为复数
			nums[index-1] = -abs(nums[index-1])
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 { // 找到都一个大于0的数的索引
			return i + 1
		}
	}

	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
