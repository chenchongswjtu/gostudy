package main

// 1644.二叉树的最近公共祖先 II
// p,q不一定在root中，不存在返回nil

var foundP, foundQ bool

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	ret := find2(root, p.Val, q.Val)
	if !foundP || !foundQ {
		return nil
	}
	return ret
}

func find2(root *TreeNode, val1 int, val2 int) *TreeNode {
	if root == nil {
		return nil
	}

	left := find2(root.Left, val1, val2)
	right := find2(root.Right, val1, val2)

	if left != nil && right != nil {
		return root
	}

	if root.Val == val1 || root.Val == val2 { // 查看p,q是否在root
		if root.Val == val1 {
			foundP = true
		}

		if root.Val == val2 {
			foundQ = true
		}
		return root
	}

	if left == nil {
		return right
	}

	return left

}
