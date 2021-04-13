package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isValidBST1(Ints2TreeNode([]int{1, 1})))
}

// 验证二叉搜索树(递归)
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

// 验证二叉搜索树(非递归，中序遍历就是一个递增数组)
func isValidBST1(root *TreeNode) bool {
	var stack []*TreeNode
	var pre = math.MinInt64 //上一次遍历的数字

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 取出最后一个
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 中序遍历就是一个递增数组,应该要大于前一个遍历的数字
		if root.Val <= pre {
			return false
		}

		pre = root.Val
		root = root.Right
	}

	return true
}
