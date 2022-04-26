package main

// 112. 路径总和
//给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，
//这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
//
//叶子节点 是指没有子节点的节点。
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil { //由于树是空的，所以不存在根节点到叶子节点的路径。
		return false
	}

	if root.Left == nil && root.Right == nil { // 叶子节点
		if root.Val == targetSum { // 相等，true
			return true
		} else { // 不相等，false
			return false
		}
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) // 递归调用
}
