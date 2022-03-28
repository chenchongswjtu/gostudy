package main

// 101. 对称二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		return isSymmetricHelper(root.Left, root.Right)
	}

	return false
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil {
		return false
	}

	if right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}
