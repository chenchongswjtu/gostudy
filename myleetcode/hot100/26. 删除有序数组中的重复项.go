package main

// 26.删除有序数组中重复项
// 快慢指针
func removeDuplicates(nums []int) int {
	n := len(nums)
	fast := 0
	slow := 0

	for fast < n {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}
