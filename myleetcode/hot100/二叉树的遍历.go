package main

import "fmt"

// 前序遍历
// 根左右
// 非递归实现
func preOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	var stack []*TreeNode
	var cur = root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			ret = append(ret, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}

	return ret
}

// 中序遍历
// 左根右
// 非递归实现
func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	var stack []*TreeNode
	var cur = root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) > 0 {
			cur = stack[len(stack)-1]
			ret = append(ret, cur.Val)
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}

	return ret
}

// 后序遍历
// 左右根
// 非递归实现
// 根右左进行reverse
func postOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	var stack []*TreeNode
	var cur = root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			ret = append(ret, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Left
		}
	}

	reverse(ret)

	return ret
}

func main() {
	fmt.Println(preOrder(Ints2TreeNode([]int{1, 2, null, 4, 5})))
	fmt.Println(inOrder(Ints2TreeNode([]int{1, 2, null, 4, 5})))
	fmt.Println(postOrder(Ints2TreeNode([]int{1, 2, null, 4, 5})))
}

// NULL 方便添加测试数据
var null = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != null {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != null {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
