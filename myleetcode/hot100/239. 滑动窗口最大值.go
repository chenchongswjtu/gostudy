package main

//239. 滑动窗口最大值
//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值。
//
//
//
//示例 1：
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7
//示例 2：
//
//输入：nums = [1], k = 1
//输出：[1]
//示例 3：
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//示例 4：
//
//输入：nums = [9,11], k = 2
//输出：[11]
//示例 5：
//
//输入：nums = [4,-2], k = 2
//输出：[4]
//
//
//提示：
//
//1 <= nums.length <= 105
//-104 <= nums[i] <= 104
//1 <= k <= nums.length

// window[0]是最大的数据的索引
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k) // store the index of nums
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums { // if the left-most index is out of window, remove it
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
			window = window[0 : len(window)-1]
		}
		window = append(window, i) // store the index of nums
		if i >= k-1 {
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums
		}
	}
	return result
}

// 单调队列（递减，最大的在第一个）
func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}

	descQueue := make([]int, 0) // index of nums
	ret := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		for len(descQueue) > 0 && nums[descQueue[len(descQueue)-1]] < nums[i] {
			descQueue = descQueue[:len(descQueue)-1] // 从末尾删除所有小于nums[i]的index
		}

		descQueue = append(descQueue, i)
		if descQueue[0] < i-k+1 {
			descQueue = descQueue[1:] // 删除第一个index，index在k的范围外，为最大的数
		}

		if i >= k-1 { // 从i为k-1开始
			ret = append(ret, nums[descQueue[0]])
		}

	}
	return ret
}
