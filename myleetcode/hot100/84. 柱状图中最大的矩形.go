package main

/*对比
11. 盛最多水的容器    64.1%    中等【双指针，单调栈不可行】
42. 接雨水	55.7%	困难【双指针+递归，动态规划，双指针，栈】
84. 柱状图中最大的矩形	42.9%	困难【单调栈（单减栈）】
*/

//84. 柱状图中最大的矩形
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//
//
//以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//
//
//
//
//
//图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//
//
//示例:
//
//输入: [2,1,5,6,2,3]
//输出: 10

// 暴力解法（超时）
func largestRectangleArea(heights []int) int {
	n := len(heights)
	var ret int

	for i := 0; i < n; i++ {
		var l int
		var r int
		for l = i; l >= 0; l-- {
			if heights[l] < heights[i] {
				break
			}
		}

		for r = i; r < n; r++ {
			if heights[r] < heights[i] {
				break
			}
		}

		ret = maxInt(ret, (r-l-1)*heights[i])
	}
	return ret
}

// 添加哨兵,heights的前后都加上0的值，
// 单调栈
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return heights[0]
	}

	var area int
	newHeights := make([]int, n+2)
	for i := 0; i < n; i++ {
		newHeights[i+1] = heights[i]
	}
	stack := []int{0}          // 将newHeights的0索引添加上
	for i := 1; i < n+2; i++ { // 从1开始索引
		for newHeights[stack[len(stack)-1]] > newHeights[i] {
			h := newHeights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i - stack[len(stack)-1] - 1 // 单调栈中newHeights[i]小于栈顶，栈顶下面也小于
			area = maxInt(area, w*h)
		}
		stack = append(stack, i)
	}
	return area
}
