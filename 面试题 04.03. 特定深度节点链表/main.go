package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type levelTreeNode struct {
	node  *TreeNode
	level int
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}

	treeNodeSlice := make([]*levelTreeNode, 0)
	treeNodeSlice = append(treeNodeSlice, &levelTreeNode{
		node:  tree,
		level: 0,
	})

	index := 0
	for index <= len(treeNodeSlice)-1 {
		temp := treeNodeSlice[index]
		index++
		if temp.node.Left != nil {
			treeNodeSlice = append(treeNodeSlice, &levelTreeNode{node: temp.node.Left, level: temp.level + 1})
		}

		if temp.node.Right != nil {
			treeNodeSlice = append(treeNodeSlice, &levelTreeNode{node: temp.node.Right, level: temp.level + 1})
		}
	}

	preLevel := treeNodeSlice[0].level

	res := make([]*ListNode, 0)
	last := make([]*ListNode, 0)
	list := &ListNode{
		Val:  treeNodeSlice[0].node.Val,
		Next: nil,
	}

	res = append(res, list)
	last = append(last, list)

	for i := 1; i < len(treeNodeSlice); i++ {
		if treeNodeSlice[i].level == preLevel {
			list := &ListNode{
				Val:  treeNodeSlice[i].node.Val,
				Next: nil,
			}

			last[preLevel].Next = list
			last[preLevel] = list
		} else {
			list := &ListNode{
				Val:  treeNodeSlice[i].node.Val,
				Next: nil,
			}
			res = append(res, list)
			last = append(last, list)
		}

		preLevel = treeNodeSlice[i].level
	}

	return res
}

func main() {

}
