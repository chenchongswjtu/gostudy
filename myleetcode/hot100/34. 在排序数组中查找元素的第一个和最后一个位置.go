package main

//34. 在排序数组中查找元素的第一个和最后一个位置
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//进阶：
//
//你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
//
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//示例 2：
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//示例 3：
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//提示：
//
//0 <= nums.length <= 105
//-109 <= nums[i] <= 109
//nums 是一个非递减数组
//-109 <= target <= 109

// 二分法
func searchRange(nums []int, target int) []int {
	return []int{searchLeftRange(nums, target), searchRightRange(nums, target)}
}

func searchLeftRange(nums []int, target int) int {
	size := len(nums)
	l, r := 0, size-1
	for {
		if l > r {
			break
		}

		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			r = m - 1 // 收缩右边界
		}
	}

	if l >= size || nums[l] != target {
		return -1
	}

	return l
}

func searchRightRange(nums []int, target int) int {
	size := len(nums)
	l, r := 0, size-1
	for {
		if l > r {
			break
		}

		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1 // 收缩左边界
		}
	}

	if r < 0 || nums[r] != target {
		return -1
	}

	return r
}
