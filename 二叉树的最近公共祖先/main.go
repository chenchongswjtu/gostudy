package main

import (
	"fmt"

	"github.com/chenchongswjtu/gostudy/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

func main() {
	root := tree.Ints2TreeNode([]int{3, 5, 1, 6, 2, 0, 8, tree.NULL, tree.NULL, 7, 4})
	p := &tree.Node{Val: 5}
	q := &tree.Node{Val: 1}
	node := lowestCommonAncestor(root, p, q)
	fmt.Println(node.Val)
}
