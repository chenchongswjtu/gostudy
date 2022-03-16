package main

/*对比
11. 盛最多水的容器    64.1%    中等【双指针，单调栈不可行】
42. 接雨水	55.7%	困难【双指针+递归，动态规划，双指针，栈】
84. 柱状图中最大的矩形	42.9%	困难【单调栈（单减栈）】
*/

//42. 接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
//示例 1：
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//示例 2：
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//提示：
//
//n == height.length
//0 <= n <= 3 * 104
//0 <= height[i] <= 105

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
