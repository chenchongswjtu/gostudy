package main

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
//叶子节点 是指没有子节点的节点。

// 113. 路径总和 II
func pathSum(root *TreeNode, targetSum int) [][]int {
	var allPath = make([][]int, 0)
	pathSumHelper(root, targetSum, []int{}, &allPath)
	return allPath
}

func pathSumHelper(root *TreeNode, sum int, path []int, allPath *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			path = append(path, root.Val)
			*allPath = append(*allPath, append([]int(nil), path...)) // 将one的路径放到一个新的空间中
		}
		return
	}

	path = append(path, root.Val)
	pathSumHelper(root.Left, sum-root.Val, path, allPath)
	pathSumHelper(root.Right, sum-root.Val, path, allPath)
}
