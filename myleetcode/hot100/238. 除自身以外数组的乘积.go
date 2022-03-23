package main

//238. 除自身以外数组的乘积
//给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
//
//
//
//示例:
//
//输入: [1,2,3,4]
//输出: [24,12,8,6]
//
//
//提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
//
//说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
//进阶：
//你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

// 使用除法，两趟遍历 + 分类讨论
func productExceptSelf(nums []int) []int {
	var total = 1
	var zeroCount int
	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}
		total *= num
	}

	var ret = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if zeroCount == 0 {
			ret[i] = total / nums[i]
		} else if zeroCount == 1 && nums[i] == 0 {
			ret[i] = total
		} else {
			ret[i] = 0
		}
	}
	return ret
}

// 左右乘积
func productExceptSelf1(nums []int) []int {
	var size = len(nums)
	var left = make([]int, size)
	var right = make([]int, size)
	var ret = make([]int, size)

	left[0] = 1
	for i := 1; i < size; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[size-1] = 1
	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < size; i++ {
		ret[i] = left[i] * right[i]
	}

	return ret
}
