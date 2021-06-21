package main

func main() {

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
