package main

// 26.移除元素
// 快慢指针
func removeElement(nums []int, val int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
