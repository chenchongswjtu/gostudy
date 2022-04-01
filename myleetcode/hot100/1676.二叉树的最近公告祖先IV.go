package main

// 依然给你输入一棵不含重复值的二叉树，但这次不是给你输入p和q两个节点了，
// 而是给你输入一个包含若干节点的列表nodes（这些节点都存在于二叉树中），让你算这些节点的最近公共祖先。
func lowestCommonAncestor4(root *TreeNode, nodes []*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	set := map[int]struct{}{}
	for _, node := range nodes {
		set[node.Val] = struct{}{}
	}

	if _, ok := set[root.Val]; ok { //检查nodes中存不存在root
		return root
	}

	left := lowestCommonAncestor4(root.Left, nodes)
	right := lowestCommonAncestor4(root.Right, nodes)
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}
