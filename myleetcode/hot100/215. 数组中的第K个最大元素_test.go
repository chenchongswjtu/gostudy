package main

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	nums := []int{1, 2, 3, 4, 5}
	shuffle(nums)
	fmt.Println(nums)

	fmt.Println(findKthLargest1([]int{3, 2, 1, 5, 6, 4}, 2))
}
