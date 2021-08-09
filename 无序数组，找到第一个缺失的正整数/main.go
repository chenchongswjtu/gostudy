package main

import (
	"fmt"
)

// [3,4,-1,1] return 2
func findFirstMissPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] <= 0 || nums[i] > n || nums[i] == i+1 {
			i++
		} else {
			if nums[i] == nums[nums[i]-1] {
				i++
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}

	for i := 0; i < n; {
		if nums[i] == i+1 {
			i++
		} else {
			return i + 1
		}
	}

	return n + 1
}
func main() {
	fmt.Println(findFirstMissPositive([]int{3, 4, -1, 1}))
	fmt.Println(findFirstMissPositive([]int{1, 2, 3, 4}))
	fmt.Println(findFirstMissPositive([]int{1, 2, 0}))
	fmt.Println(findFirstMissPositive([]int{7, 8, 9, 11, 12}))
}
