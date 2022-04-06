package main

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	v1 := minInt(p.Val, q.Val)
	v2 := maxInt(p.Val, q.Val)
	return find3(root, v1, v2)
}

func find3(root *TreeNode, v1 int, v2 int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > v2 {
		return find3(root.Left, v1, v2)
	}

	if root.Val < v1 {
		return find3(root.Right, v1, v2)
	}

	return root
}
