package main

// 103. 二叉树的锯齿形层序遍历

//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		tmp := append([]*TreeNode(nil), q...)
		q = q[:0]
		var t []int
		for _, node := range tmp {
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			t = append(t, node.Val)
		}
		result = append(result, t)
	}

	for i, ints := range result {
		if i%2 == 1 {
			reverse(ints)
		}
	}
	return result
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
