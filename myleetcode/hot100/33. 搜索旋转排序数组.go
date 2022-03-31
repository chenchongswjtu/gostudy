package main

// 33. 搜索旋转排序数组
// 二分法
func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for {
		if left > right {
			break
		}

		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[left] == target {
			return left
		} else if nums[right] == target {
			return right
		}

		if nums[mid] > nums[left] {
			if nums[left] < target && target < nums[mid] {
				right = mid - 1
			} else if nums[mid] < target || target < nums[right] {
				left = mid + 1
			} else {
				return -1
			}
		} else {
			if nums[mid] < target && target < nums[right] {
				left = mid + 1
			} else if nums[left] < target || target < nums[mid] {
				right = mid - 1
			} else {
				return -1
			}
		}
	}
	return -1
}
