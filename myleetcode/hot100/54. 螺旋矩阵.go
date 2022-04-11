package main

//54. 螺旋矩阵
//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 层遍历
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m := len(matrix)
	n := len(matrix[0])
	total := m * n
	order := make([]int, total)
	index := 0
	left, right, top, bottom := 0, n-1, 0, m-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			order[index] = matrix[top][i]
			index++
		}

		for i := top + 1; i <= bottom; i++ {
			order[index] = matrix[i][right]
			index++
		}

		if left < right && top < bottom {
			for i := right - 1; i >= left; i-- {
				order[index] = matrix[bottom][i]
				index++
			}

			for i := bottom - 1; i >= top+1; i-- {
				order[index] = matrix[i][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return order
}
