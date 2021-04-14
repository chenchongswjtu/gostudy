package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(isValidBST1(Ints2TreeNode([]int{1, 1})))
	fmt.Println(levelOrder(Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7})))
}

// 验证二叉搜索树(递归)
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

// 验证二叉搜索树(非递归，中序遍历就是一个递增数组)
func isValidBST1(root *TreeNode) bool {
	var stack []*TreeNode
	var pre = math.MinInt64 //上一次遍历的数字

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 取出最后一个
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 中序遍历就是一个递增数组,应该要大于前一个遍历的数字
		if root.Val <= pre {
			return false
		}

		pre = root.Val
		root = root.Right
	}

	return true
}

// 中序遍历（递归）
func inorderTraversal(root *TreeNode) []int {
	var res []int
	help1(root, &res)
	return res
}

func help1(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	help1(root.Left, res)
	*res = append(*res, root.Val)
	help1(root.Right, res)
}

// 中序遍历（非递归）
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 具体操作
		res = append(res, root.Val)

		root = root.Right
	}

	return res
}

// 判断两个树是否相同
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil {
		return false
	}

	if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 是否是对称二叉树
func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelper(root, root)
}

func isSymmetricHelper(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return false
	}

	if b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isSymmetricHelper(a.Left, b.Right) && isSymmetricHelper(a.Right, b.Left)
}

type levelNode struct {
	node  *TreeNode
	level int
}

// 二叉树的层序
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	var queue []*levelNode
	var index int
	queue = append(queue, &levelNode{
		node:  root,
		level: 0,
	})
	for index < len(queue) {
		temp := queue[index]
		index++
		if temp.node.Left != nil {
			queue = append(queue, &levelNode{
				node:  temp.node.Left,
				level: temp.level + 1,
			})
		}
		if temp.node.Right != nil {
			queue = append(queue, &levelNode{
				node:  temp.node.Right,
				level: temp.level + 1,
			})
		}
	}

	res = append(res, []int{queue[0].node.Val})
	preLevel := 0
	for i := 1; i < len(queue); i++ {
		if queue[i].level != preLevel {
			res = append(res, []int{queue[i].node.Val})
		} else {
			res[queue[i].level] = append(res[queue[i].level], queue[i].node.Val)
		}
		preLevel = queue[i].level
	}

	return res
}

//type levelNode struct {
//	node  *TreeNode
//	level int
//}

// 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	var queue []*levelNode
	var index int
	queue = append(queue, &levelNode{
		node:  root,
		level: 0,
	})
	for index < len(queue) {
		temp := queue[index]
		index++
		if temp.node.Left != nil {
			queue = append(queue, &levelNode{
				node:  temp.node.Left,
				level: temp.level + 1,
			})
		}
		if temp.node.Right != nil {
			queue = append(queue, &levelNode{
				node:  temp.node.Right,
				level: temp.level + 1,
			})
		}
	}

	res = append(res, []int{queue[0].node.Val})
	preLevel := 0
	for i := 1; i < len(queue); i++ {
		if queue[i].level != preLevel {
			res = append(res, []int{queue[i].node.Val})
		} else {
			res[queue[i].level] = append(res[queue[i].level], queue[i].node.Val)
		}
		preLevel = queue[i].level
	}

	for i, v := range res {
		if i%2 == 0 {
			res[i] = v
		} else {
			res[i] = reverse(v)
		}
	}

	return res
}

func reverse(a []int) []int {
	i := 0
	j := len(a) - 1

	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	return a
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 105. 从前序与中序遍历序列构造二叉树
// 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func buildTreeHelper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	if inStart > inEnd {
		return nil
	}

	rootVal := preorder[preStart]
	var pos int

	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			pos = i
		}
	}

	l := pos - inStart

	root := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}

	root.Left = buildTreeHelper(preorder, preStart+1, preStart+l, inorder, inStart, pos-1)
	root.Right = buildTreeHelper(preorder, preStart+l+1, preEnd, inorder, pos+1, inEnd)

	return root
}

// 106. 从中序与后序遍历序列构造二叉树
func buildTree1(inorder []int, postorder []int) *TreeNode {
	return buildTreeHelper1(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildTreeHelper1(inorder []int, s1 int, e1 int, postorder []int, s2 int, e2 int) *TreeNode {
	if s1 > e1 || s2 > e2 {
		return nil
	}

	rootVal := postorder[e2]
	var pos int

	for i := s1; i <= e1; i++ {
		if inorder[i] == rootVal {
			pos = i
		}
	}

	l := pos - s1

	root := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}

	root.Left = buildTreeHelper1(inorder, s1, pos-1, postorder, s2, s2+l-1)
	root.Right = buildTreeHelper1(inorder, pos+1, e1, postorder, s2+l, e2-1)

	return root
}
