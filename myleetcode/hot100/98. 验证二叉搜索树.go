package main

import "math"

//98. 验证二叉搜索树
//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//示例 1:
//
//输入:
//2
/// \
//1   3
//输出: true
//示例 2:
//
//输入:
//5
/// \
//1   4
/// \
//3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//根节点的值为 5 ，但是其右子节点值为 4 。

// 递归实现
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		return root.Val > maxValBST(root.Left) && root.Val < minValBST(root.Right) && isValidBST(root.Left) && isValidBST(root.Right)
	}

	if root.Left == nil && root.Right != nil {
		return root.Val < minValBST(root.Right) && isValidBST(root.Right)
	}

	return root.Val > maxValBST(root.Left) && isValidBST(root.Left)
}

func minValBST(node *TreeNode) int {
	for node.Left != nil {
		node = node.Left
	}

	return node.Val
}

func maxValBST(node *TreeNode) int {
	for node.Right != nil {
		node = node.Right
	}

	return node.Val
}

//中序遍历
func isValidBST1(root *TreeNode) bool {
	var preVal = math.MinInt64
	var helper func(root *TreeNode) bool
	helper = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !helper(root.Left) {
			return false
		}

		if root.Val > preVal {
			preVal = root.Val
		} else {
			return false
		}

		if !helper(root.Right) {
			return false
		}

		return true
	}
	return helper(root)
}
