package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

// 215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	low, high := 0, len(nums)-1
	target := len(nums) - k // 数组索引下标

	for {
		mid := partition(nums, low, high)
		if mid == target {
			return nums[mid]
		} else if mid < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}

func partition(nums []int, low int, high int) int {
	t := nums[low]
	for low < high {
		for low < high && nums[high] >= t {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= t {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = t
	return low
}
