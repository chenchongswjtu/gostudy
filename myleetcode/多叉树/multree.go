package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 429.多叉树的层序遍历
func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	var queue = []*Node{root}
	for len(queue) > 0 {
		one := queue
		queue = nil
		v := make([]int, 0)
		for _, o := range one {
			v = append(v, o.Val)
			for _, n := range o.Children {
				if n != nil {
					queue = append(queue, n)
				}
			}
		}
		res = append(res, v)
	}
	return res
}

// 559. N 叉树的最大深度
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	maxD := 0
	for _, c := range root.Children {
		maxD = max(maxD, maxDepth(c))
	}

	return maxD + 1
}

// 589. N 叉树的前序遍历
func preorder(root *Node) []int {
	var ans []int
	preorderHelper(root, &ans)
	return ans
}

func preorderHelper(root *Node, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	for _, c := range root.Children {
		preorderHelper(c, ans)
	}
}

// 590. N 叉树的后序遍历
func postorder(root *Node) []int {
	var ans []int
	postorderHelper(root, &ans)
	return ans
}

func postorderHelper(root *Node, ans *[]int) {
	if root == nil {
		return
	}

	for _, c := range root.Children {
		postorderHelper(c, ans)
	}
	*ans = append(*ans, root.Val)
}
