package main

//85. 最大矩形
//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
//
//
//示例 1：
//
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
//输出：6
//解释：最大矩形如上图所示。
//示例 2：
//
//输入：matrix = []
//输出：0
//示例 3：
//
//输入：matrix = [["0"]]
//输出：0
//示例 4：
//
//输入：matrix = [["1"]]
//输出：1
//示例 5：
//
//输入：matrix = [["0","0"]]
//输出：0
//
//
//提示：
//
//rows == matrix.length
//cols == matrix[0].length
//0 <= row, cols <= 200
//matrix[i][j] 为 '0' 或 '1'

// 每一层可以看作是柱状图中的最大矩形
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	ret := 0
	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		ret = maxInt(ret, largestRectangleArea2(heights))
	}
	return ret
}
