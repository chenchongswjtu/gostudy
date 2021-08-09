package main

import "fmt"

func exchange(nums []int) []int {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4, 5}))
}
