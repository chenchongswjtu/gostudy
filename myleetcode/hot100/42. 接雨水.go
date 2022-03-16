package main

func trap(height []int) int {
	size := len(height)
	if size <= 2 {
		return 0
	}

	// 维护两个数组，分别记录每根柱子左边最大和右边最大值
	leftMax := make([]int, size)
	rightMax := make([]int, size)

	leftMax[0] = height[0]
	rightMax[size-1] = height[size-1]

	for i := 1; i < size; i++ {
		leftMax[i] = maxInt(height[i], leftMax[i-1])
	}

	for i := size - 2; i >= 0; i-- {
		rightMax[i] = maxInt(height[i], rightMax[i+1])
	}

	sum := 0
	for i := 0; i < size; i++ {
		sum += minInt(leftMax[i], rightMax[i]) - height[i]
	}
	return sum
}
