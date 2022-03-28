package main

import "math"

//124. 二叉树中的最大路径和
//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
//路径和 是路径中各节点值的总和。
//
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//
//
//示例 1：
//
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//示例 2：
//
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
//提示：
//
//树中节点数目范围是 [1, 3 * 104]
//-1000 <= Node.val <= 1000

func maxPathSum(root *TreeNode) int {
	var ret = math.MinInt64
	postOrderer(root, &ret)
	return ret
}

// 包括该节点的最大值
func postOrderer(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := maxInt(postOrderer(root.Left, max), 0)
	right := maxInt(postOrderer(root.Right, max), 0)

	*max = maxInt(*max, left+right+root.Val)

	// 只能将其中一个返回上去
	return root.Val + maxInt(left, right)
}
