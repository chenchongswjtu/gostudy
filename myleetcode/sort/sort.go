package main

import "fmt"

func main() {
	var nums = []int{3, 2, 41, 4, 2}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// 快速排序
func quickSort(nums []int, low int, high int) {
	if low >= high {
		return
	}
	mid := partition(nums, low, high)
	quickSort(nums, low, mid-1)
	quickSort(nums, mid+1, high)
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
