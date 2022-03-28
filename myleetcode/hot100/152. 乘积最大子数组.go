package main

//152. 乘积最大子数组
//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//
//
//示例 1:
//
//输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//
//输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

// 动态规划（空间优化：滚动数组）
func maxProduct(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	dpMin, dpMax := nums[0], nums[0]
	ret := nums[0]
	for i := 1; i < size; i++ {
		tmpMin, tmpMax := dpMin, dpMax
		if nums[i] < 0 {
			tmpMin = minInt(nums[i], nums[i]*dpMax)
			tmpMax = maxInt(nums[i], nums[i]*dpMin)
		} else if nums[i] > 0 {
			tmpMin = minInt(nums[i], nums[i]*dpMin)
			tmpMax = maxInt(nums[i], nums[i]*dpMax)
		} else {
			tmpMin, tmpMax = 0, 0
		}
		dpMin, dpMax = tmpMin, tmpMax
		ret = maxInt(ret, dpMax)
	}
	return ret
}
