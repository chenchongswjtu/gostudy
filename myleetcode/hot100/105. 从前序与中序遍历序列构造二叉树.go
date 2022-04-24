package main

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	v := preorder[0]
	var i int
	for i = 0; i < len(inorder); i++ {
		if v == inorder[i] {
			break
		}
	}

	root := &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}

	root.Left = buildTree(preorder[1:1+i], inorder[:i])
	root.Right = buildTree(preorder[1+i:], inorder[i+1:])
	return root
}
