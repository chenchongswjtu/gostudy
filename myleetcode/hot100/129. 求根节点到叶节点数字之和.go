package main

// 129. 求根节点到叶节点数字之和
//给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//每条从根节点到叶节点的路径都代表一个数字：
//
//例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
//计算从根节点到叶节点生成的 所有数字之和 。
//
//叶节点 是指没有子节点的节点。
// 使用深度搜索
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum = 0
	var dfs = func(node *TreeNode, preSum int) {}
	dfs = func(node *TreeNode, preSum int) {
		if node.Left == nil && node.Right == nil {
			sum += preSum*10 + node.Val
			return
		}

		if node.Left != nil {
			dfs(node.Left, preSum*10+node.Val)
		}

		if node.Right != nil {
			dfs(node.Right, preSum*10+node.Val)
		}
	}

	dfs(root, 0)
	return sum
}
