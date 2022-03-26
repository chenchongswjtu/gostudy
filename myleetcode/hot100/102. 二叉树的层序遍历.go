package main

//102. 二叉树的层序遍历
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其层序遍历结果：
//
//[
//[3],
//[9,20],
//[15,7]
//]

// 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		tmp := make([]*TreeNode, len(stack)) // 重新copy一份
		copy(tmp, stack)
		stack = stack[0:0]

		var nums = make([]int, len(tmp))
		for i, node := range tmp {
			nums[i] = node.Val
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		ret = append(ret, nums)
	}
	return ret
}
