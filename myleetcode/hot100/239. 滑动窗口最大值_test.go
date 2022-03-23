package main

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow1([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
