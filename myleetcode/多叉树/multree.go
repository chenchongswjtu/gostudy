package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
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
