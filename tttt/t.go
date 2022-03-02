package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subSet([]int{1, 2, 3}))
	fmt.Println(myPow(2.00, -10))
}

func subSet(nums []int) [][]int {
	all := make([][]int, 0)
	one := make([]int, 0)
	subSetHelper(nums, 0, one, &all)
	return all
}

func subSetHelper(nums []int, i int, one []int, all *[][]int) {
	if i == len(nums) {
		*all = append(*all, one)
		return
	}

	// 选择nums[i]
	one = append(one, nums[i])
	subSetHelper(nums, i+1, one, all)
	// 不选择nums[i]
	one = one[:len(one)-1]
	subSetHelper(nums, i+1, one, all)
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}

	p.Right = right
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var memo = make(map[string]int)
	var res = make([]*TreeNode, 0)
	traverse(root, memo, &res)
	return res
}

func traverse(root *TreeNode, memo map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traverse(root.Left, memo, res)
	right := traverse(root.Right, memo, res)

	subTree := left + "," + right + "," + strconv.Itoa(root.Val)

	n, ok := memo[subTree]
	if ok {
		if n == 1 {
			*res = append(*res, root)
		}
		memo[subTree]++
	} else {
		memo[subTree] = 1
	}

	return subTree
}
