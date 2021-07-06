package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB"))
}

// 剑指 Offer 04. 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	i, j := 0, n-1
	for i >= 0 && i <= m-1 && j >= 0 && j <= n-1 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 剑指 Offer 06. 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	var cur = head
	var ans []int
	for cur != nil {
		ans = append([]int{cur.Val}, ans...)
		cur = cur.Next
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指 Offer 07. 重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= 1000000007
	}

	return dp[n]
}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		sum := (a + b) % 1000000007
		a = b
		b = sum
	}

	return b
}

// 剑指 Offer 11. 旋转数组的最小数字 (二分查找)
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

// 剑指 Offer 12. 矩阵中的路径
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	if len(board[0]) == 0 {
		return false
	}

	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j, index int, visited [][]bool) bool
	dfs = func(i, j, index int, visited [][]bool) bool {
		if index == len(word) {
			return true
		}

		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}

		if visited[i][j] || board[i][j] != word[index] {
			return false
		}

		visited[i][j] = true
		res := dfs(i+1, j, index+1, visited) || dfs(i, j+1, index+1, visited) || dfs(i-1, j, index+1, visited) || dfs(i, j-1, index+1, visited)
		visited[i][j] = false // 回溯
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

// 剑指 Offer 13. 机器人的运动范围
// 广度优先搜索
func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}

	get := func(x int) int {
		res := 0
		for x != 0 {
			res += x % 10
			x = x / 10
		}
		return res
	}

	queue := make([][2]int, 0)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	// 向下和向右的方向数组
	dx := [2]int{0, 1}
	dy := [2]int{1, 0}

	queue = append(queue, [2]int{0, 0})
	visited[0][0] = true
	ans := 1
	for len(queue) > 0 {
		tail := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		x := tail[0]
		y := tail[1]
		for i := 0; i < 2; i++ {
			tx := dx[i] + x
			ty := dy[i] + y
			if tx < 0 || tx >= m || ty < 0 || ty >= n || visited[tx][ty] || get(tx)+get(ty) > k {
				continue
			}
			queue = append(queue, [2]int{tx, ty})
			visited[tx][ty] = true
			ans++
		}
	}

	return ans
}
