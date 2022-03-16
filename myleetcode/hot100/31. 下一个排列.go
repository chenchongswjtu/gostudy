package main

//31. 下一个排列
//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须 原地 修改，只允许使用额外常数空间。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[1,3,2]
//示例 2：
//
//输入：nums = [3,2,1]
//输出：[1,2,3]
//示例 3：
//
//输入：nums = [1,1,5]
//输出：[1,5,1]
//示例 4：
//
//输入：nums = [1]
//输出：[1]
//
//
//提示：
//
//1 <= nums.length <= 100
//0 <= nums[i] <= 100

// 最后升序位 交换 最后「大数」，升序位后翻转
func nextPermutation(nums []int) []int {
	//先找出最大的索引 k 满足 nums[k] < nums[k+1]，如果不存在，就翻转整个数组；
	//再找出另一个最大索引 l 满足 nums[l] > nums[k]；
	//交换 nums[l] 和 nums[k]；
	//最后翻转 nums[k+1:]。
	size := len(nums)
	k := -1
	for k = size - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}

	if k >= 0 {
		for l := size - 1; l >= 0; l-- {
			if l != k && nums[l] > nums[k] {
				nums[l], nums[k] = nums[k], nums[l]
				break
			}
		}
	}
	reverseSlice(nums[k+1:])
	return nums
}

func reverseSlice(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
