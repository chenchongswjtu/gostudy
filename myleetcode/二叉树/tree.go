package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	k := Constructor3(1, []int{-2})
	fmt.Println(k.Add(-3))
	fmt.Println(k.Add(0))
	fmt.Println(k.Add(2))
	fmt.Println(k.Add(-1))
	fmt.Println(k.Add(4))

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
	if preStart > preEnd || inStart > inEnd {
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

func buildTreeHelper1(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
	if inStart > inEnd || postStart > postEnd {
		return nil
	}

	rootVal := postorder[postEnd]
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

	root.Left = buildTreeHelper1(inorder, inStart, pos-1, postorder, postStart, postStart+l-1)
	root.Right = buildTreeHelper1(inorder, pos+1, inEnd, postorder, postStart+l, postEnd-1)

	return root
}

//type levelNode struct {
//	node  *TreeNode
//	level int
//}

// 107.二叉树的层序(从下到上)
func levelOrderBottom(root *TreeNode) [][]int {
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

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

// 108. 将有序数组转化为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (end-start)/2 + start
	root := &TreeNode{Val: nums[mid], Left: nil, Right: nil}

	root.Left = sortedArrayToBSTHelper(nums, start, mid-1)
	root.Right = sortedArrayToBSTHelper(nums, mid+1, end)

	return root
}

// 110.判断是不是平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if math.Abs(float64(maxDepth(root.Left)-maxDepth(root.Right))) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 最小高度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil && root.Left != nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//
//	minD := math.MaxInt64
//	if root.Left != nil {
//		minD = min(minD, minDepth(root.Left))
//	}
//
//	if root.Right != nil {
//		minD = min(minD, minDepth(root.Right))
//	}
//
//	return minD + 1
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 112. 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			return true
		}

		return false
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 112. 路径总和2,返回路径
func pathSum(root *TreeNode, targetSum int) [][]int {
	var allPath [][]int
	var path []int
	pathSumHelper(root, targetSum, path, &allPath)
	return allPath
}

func pathSumHelper(root *TreeNode, targetSum int, path []int, allPath *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			path = append(path, root.Val)
			// 这里将path拷贝到新的[]int,防止之后的将path覆盖
			var path1 = make([]int, len(path))
			copy(path1, path)
			*allPath = append(*allPath, path1)
			return
		}
		return
	}

	path = append(path, root.Val)
	pathSumHelper(root.Left, targetSum-root.Val, path, allPath)
	pathSumHelper(root.Right, targetSum-root.Val, path, allPath)
}

// 114.二叉树展开为链表
func flatten(root *TreeNode) {
	pre := make([]*TreeNode, 0)
	pre = preOrder(root)
	if len(pre) > 0 {
		root = pre[0]
		cur := root

		for i := 1; i < len(pre); i++ {
			cur.Left = nil
			cur.Right = pre[i]
			cur = cur.Right
		}
	}

}

func preOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	pre := make([]*TreeNode, 0)
	pre = append(pre, root)
	pre = append(pre, preOrder(root.Left)...)
	pre = append(pre, preOrder(root.Right)...)

	return pre
}

// 99. 恢复二叉搜索树 先中序遍历，找到可能乱序的节点，再交互
func recoverTree(root *TreeNode) {
	res := inOrder(root)
	if len(res) <= 1 {
		return
	}
	var x, y *TreeNode
	for i := 1; i < len(res); i++ {
		if res[i].Val < res[i-1].Val {
			y = res[i]
			if x == nil {
				x = res[i-1]
			}
		}

	}

	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}
}

func inOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	res := make([]*TreeNode, 0)
	res = append(res, inOrder(root.Left)...)
	res = append(res, root)
	res = append(res, inOrder(root.Right)...)

	return res
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116. 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	index := 0
	for index < len(queue) {
		temp := queue[index]
		index++
		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}
	}

	for i := 0; i < len(queue); i++ {
		if is2pow(i + 2) {
			queue[i].Next = nil
		} else {
			if i+1 < len(queue) {
				queue[i].Next = queue[i+1]
			}
		}
	}

	return root
}

func is2pow(n int) bool {
	base := 2
	for {
		if base == n {
			return true
		} else if base > n {
			return false
		}
		base = base * 2
	}
}

// 117. 填充每个节点的下一个右侧节点指针 II(同样适用于106)
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		oneLayer := queue
		queue = nil

		for i, node := range oneLayer {
			if i < len(oneLayer)-1 {
				node.Next = oneLayer[i+1]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

//////////////////////////////////////////////////////////////////////////
// 129. 求根节点到叶节点数字之和
func sumNumbers(root *TreeNode) int {
	allPaths := make([][]int, 0)
	path := make([]int, 0)
	findAllPaths(root, path, &allPaths)
	fmt.Println(allPaths)
	sum := 0
	for _, p := range allPaths {
		sum += pathSum1(p)
	}
	return sum
}

func findAllPaths(root *TreeNode, path []int, allPaths *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		path = append(path, root.Val)
		*allPaths = append(*allPaths, path)
		return
	}

	path = append(path, root.Val)
	if root.Left != nil {
		//path1 := make([]int, len(path))
		//copy(path1, path)
		findAllPaths(root.Left, path, allPaths)
	}

	if root.Right != nil {
		//path1 := make([]int, len(path))
		//copy(path1, path)
		findAllPaths(root.Right, path, allPaths)
	}
}

func pathSum1(path []int) int {
	sum := 0
	for _, v := range path {
		sum = sum*10 + v
	}

	return sum
}

////////////////////////////////////////////////////////////////////
//144. 二叉树的前序遍历(非递归)
func preorderTraversal(root *TreeNode) []int {
	pre := make([]int, 0)
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			pre = append(pre, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node = node.Right
	}

	return pre
}

// 二叉树的中序遍历(非递归)
func inorderTraversal2(root *TreeNode) []int {
	var in []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		in = append(in, node.Val)

		node = node.Right
	}

	return in
}

// 二叉树的后序遍历(非递归)（较难）
func postorderTraversal(root *TreeNode) (res []int) {
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Right == nil || root.Right == prev { // 判断操作
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}

// 199. 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var queue = []*TreeNode{root}
	var res []int

	for len(queue) > 0 {
		one := queue
		queue = nil
		res = append(res, one[len(one)-1].Val)

		for _, v := range one {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
	}

	return res
}

// 222. 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1

}

// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	right := root.Right

	root.Left = invertTree(right)
	root.Right = invertTree(left)

	return root
}

func kthSmallest(root *TreeNode, k int) int {
	var res []int
	kthSmallestHelper(root, &res)
	if k <= len(res) {
		return res[k-1]
	}
	return -1
}

func kthSmallestHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	kthSmallestHelper(root.Left, res)
	*res = append(*res, root.Val)
	kthSmallestHelper(root.Right, res)
}

// 235. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}

/////////////////////////////////////////////////////////////////
// 236. 二叉树的最近公共祖先
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

/////////////////////////////////////////////////////
// 257. 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var paths []string
	binaryTreePathsHelper(root, "", &paths)
	return paths
}

func binaryTreePathsHelper(root *TreeNode, path string, paths *[]string) {
	if len(path) == 0 {
		path = path + strconv.Itoa(root.Val)
	} else {
		path = path + "->" + strconv.Itoa(root.Val)
	}

	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
		return
	}

	if root.Left != nil {
		binaryTreePathsHelper(root.Left, path, paths)
	}

	if root.Right != nil {
		binaryTreePathsHelper(root.Right, path, paths)
	}
}

///////////////////////////////////////////////
// 337. 打家劫舍 III
var memo = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if v, ok := memo[root]; ok {
		return v
	}

	left := 0
	right := 0
	if root.Left != nil {
		left = rob(root.Left.Left) + rob(root.Left.Right)
	} else {
		left = 0
	}

	if root.Right != nil {
		right = rob(root.Right.Left) + rob(root.Right.Right)
	} else {
		right = 0
	}

	do := root.Val + left + right
	notDo := rob(root.Left) + rob(root.Right)
	res := max(do, notDo)

	memo[root] = res

	return res
}

// 左叶子和
func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesHelper(root, false)
}

func sumOfLeftLeavesHelper(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}

	if root.Left != nil && root.Right == nil {
		return sumOfLeftLeavesHelper(root.Left, true)
	}

	if root.Left == nil && root.Right != nil {
		return sumOfLeftLeavesHelper(root.Right, false)
	}

	return sumOfLeftLeavesHelper(root.Left, true) + sumOfLeftLeavesHelper(root.Right, false)
}

/////////////////////////////////////////////
// 437. 路径总和 III
func pathSum3(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	sum := 0
	sum += pathSumStartWithRoot(root, targetSum) + pathSum3(root.Left, targetSum) + pathSum3(root.Right, targetSum)

	return sum
}

func pathSumStartWithRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	var count int

	if root.Val == sum {
		count++
	}

	leftCount := pathSumStartWithRoot(root.Left, sum-root.Val)
	rightCount := pathSumStartWithRoot(root.Right, sum-root.Val)
	count += leftCount + rightCount
	return count
}

////////////////////////////////////////
// 450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

// root的后驱
func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}

	return root.Val
}

// root的前驱
func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}

	return root.Val
}

func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		oneLayer := queue
		queue = nil
		res = oneLayer[0].Val
		for _, node := range oneLayer {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}

// 515. 在每个树行中找最大值
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var res []int
	for len(queue) > 0 {
		oneLayer := queue
		queue = nil
		max := oneLayer[0].Val
		for _, node := range oneLayer {
			if node.Val > max {
				max = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, max)
	}

	return res
}

// 530. 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res []int
	dfs(root, &res)
	if len(res) < 2 {
		return 0
	}

	min := math.MaxInt64
	for i := 1; i < len(res); i++ {
		if res[i]-res[i-1] < min {
			min = res[i] - res[i-1]
		}
	}

	return min
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}

// 543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	w := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := dfs(root.Left)
		r := dfs(root.Right)

		w = max(w, l+r)
		return 1 + max(l, r)
	}

	dfs(root)

	return w
}

func findTilt(root *TreeNode) int {
	tilt := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := dfs(root.Left)
		r := dfs(root.Right)

		tilt += int(math.Abs(float64(l - r)))
		return l + r + root.Val
	}

	dfs(root)

	return tilt
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}

	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}

	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")(" + tree2str(t.Right) + ")"
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root := &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  nil,
		Right: nil,
	}

	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)

	return root
}

// 623. 在二叉树中增加一行
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		n := &TreeNode{Val: val}
		n.Left = root
		return n
	}

	insert(root, val, 1, depth)
	return root
}

func insert(node *TreeNode, val int, n int, depth int) {
	if node == nil {
		return
	}

	if n == depth-1 {
		l := node.Left
		node.Left = &TreeNode{Val: val}
		node.Left.Left = l

		r := node.Right
		node.Right = &TreeNode{Val: val}
		node.Right.Right = r
	} else {
		insert(node.Left, val, n+1, depth)
		insert(node.Right, val, n+1, depth)
	}
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	res := make([]float64, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		o := q
		q = nil
		var sum float64
		for _, t := range o {
			if t.Left != nil {
				q = append(q, t.Left)
			}

			if t.Right != nil {
				q = append(q, t.Right)
			}
			sum += float64(t.Val)
		}

		res = append(res, sum/float64(len(o)))
	}

	return res
}

// 两数之和
func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]struct{})
	return findTargetHelper(root, k, set)
}

func findTargetHelper(root *TreeNode, k int, set map[int]struct{}) bool {
	if root == nil {
		return false
	}

	if _, ok := set[k-root.Val]; ok {
		return true
	}

	set[root.Val] = struct{}{}

	return findTargetHelper(root.Left, k, set) || findTargetHelper(root.Right, k, set)
}

// 查找重复的子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	var count = make(map[string]int)
	findDuplicateSubtreesHelper(root, count, &res)
	return res
}

func findDuplicateSubtreesHelper(node *TreeNode, count map[string]int, res *[]*TreeNode) string {
	if node == nil {
		return "#"
	}

	s := strconv.Itoa(node.Val) + "," + findDuplicateSubtreesHelper(node.Left, count, res) + "," + findDuplicateSubtreesHelper(node.Right, count, res)

	count[s]++

	if count[s] == 2 {
		*res = append(*res, node)
	}

	return s
}

// 654.最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return constructMaximumBinaryTreeHelper(nums, 0, len(nums)-1)
}

func constructMaximumBinaryTreeHelper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	max := nums[start]
	maxIndex := start
	for i := start; i <= end; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}

	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTreeHelper(nums, start, maxIndex-1)
	root.Right = constructMaximumBinaryTreeHelper(nums, maxIndex+1, end)
	return root
}

// 655. 输出二叉树
func printTree(root *TreeNode) [][]string {
	h := getHeight(root)
	res := make([][]string, h)

	for i := 0; i < h; i++ {
		res[i] = make([]string, (1<<h)-1)
	}
	fill(res, root, 0, 0, len(res[0]))

	return res
}

func fill(res [][]string, root *TreeNode, i int, l int, r int) {
	if root == nil {
		return
	}
	res[i][(l+r)/2] = strconv.Itoa(root.Val)
	fill(res, root.Left, i+1, l, (l+r)/2)
	fill(res, root.Right, i+1, (l+r+1)/2, r)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

// 662. 二叉树最大宽度
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type node struct {
		n *TreeNode
		i int
	}

	res := make([][]*node, 0)
	q := []*node{{root, 0}}
	for len(q) > 0 {
		o := q
		res = append(res, o)
		q = nil
		for _, t := range o {
			if t.n.Left != nil {
				q = append(q, &node{t.n.Left, 2 * t.i})
			}

			if t.n.Right != nil {
				q = append(q, &node{t.n.Right, 2*t.i + 1})
			}
		}
	}

	max := 0
	for i := 0; i < len(res); i++ {
		n := len(res[i])
		if max < (res[i][n-1].i - res[i][0].i) {
			max = res[i][n-1].i - res[i][0].i
		}
	}

	return max + 1
}

// 669. 修剪二叉搜索树
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	if root.Val == low {
		root.Left = nil
		root.Right = trimBST(root.Right, low, high)
		return root
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	if root.Val == high {
		root.Right = nil
		root.Left = trimBST(root.Left, low, high)
		return root
	}

	if root.Left != nil {
		root.Left = trimBST(root.Left, low, root.Val-1)
	}

	if root.Right != nil {
		root.Right = trimBST(root.Right, root.Val+1, high)
	}

	return root
}

// 671. 二叉树中第二小的节点
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	return findSecondMinimumValueHelp(root, root.Val)
}

func findSecondMinimumValueHelp(root *TreeNode, v int) int {
	if root == nil {
		return -1
	}

	if root.Val > v {
		return root.Val
	}

	left := findSecondMinimumValueHelp(root.Left, v)
	right := findSecondMinimumValueHelp(root.Right, v)

	if left == -1 && right == -1 {
		return -1
	}

	if left == -1 {
		return right
	}

	if right == -1 {
		return left
	}

	return min(left, right)
}

// 684. 冗余连接(并查集)
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	// 默认为自己连通自己
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(from, to int) bool {
		x, y := find(from), find(to)
		// 已经连通，返回false
		if x == y {
			return false
		}
		// 将x,y连通
		parent[x] = y
		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}

	return nil
}

// 687. 最长同值路径
func longestUnivaluePath(root *TreeNode) int {
	var ans int
	longestUnivaluePathHelper(root, &ans)
	return ans
}

func longestUnivaluePathHelper(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}

	l := longestUnivaluePathHelper(node.Left, ans)
	r := longestUnivaluePathHelper(node.Right, ans)

	l1, r1 := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		l1 += l + 1
	}

	if node.Right != nil && node.Right.Val == node.Val {
		r1 += r + 1
	}

	*ans = max(*ans, l1+r1)

	return max(l1, r1)
}

// 700. 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}

	return searchBST(root.Left, val)
}

// 701. 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val, Left: nil, Right: nil}
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

// 783. 二叉搜索树节点最小距离
func minDiffInBST(root *TreeNode) int {
	var res []int
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}

	inOrder(root)

	if len(res) < 2 {
		return 0
	}

	ans := res[1] - res[0]
	for i := 2; i < len(res); i++ {
		ans = min(ans, res[i]-res[i-1])
	}

	return ans
}

// 814. 二叉树剪枝
func pruneTree(root *TreeNode) *TreeNode {
	if !hasOne(root) {
		return nil
	}

	if !hasOne(root.Left) {
		root.Left = nil
	} else {
		root.Left = pruneTree(root.Left)
	}

	if !hasOne(root.Right) {
		root.Right = nil
	} else {
		root.Right = pruneTree(root.Right)
	}
	return root
}

func hasOne(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == 1 {
		return true
	}

	return hasOne(root.Left) || hasOne(root.Right)
}

// 889. 根据前序和后序遍历构造二叉树
//输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
//输出：[1,2,3,4,5,6,7]
func constructFromPrePost(pre []int, post []int) *TreeNode {
	n := len(pre)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   pre[0],
		Left:  nil,
		Right: nil,
	}

	if n == 1 {
		return root
	}

	length := 0
	for i := 0; i < n; i++ {
		if post[i] == pre[1] {
			length = i
			break
		}
	}

	root.Left = constructFromPrePost(pre[1:1+length+1], post[:length+1])
	root.Right = constructFromPrePost(pre[1+length+1:n], post[length+1:n-1])

	return root
}

// 124. 二叉树中的最大路径和(困难)
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		newGain := node.Val + leftGain + rightGain
		// 更新答案
		maxSum = max(maxSum, newGain)
		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

// 297. 二叉树的序列化与反序列化(困难)
type Codec struct{}

func Constructor() (_ Codec) {
	return
}

func (Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	s := sb.String()
	return s[:len(s)-1]
}

func (Codec) deserialize(data string) *TreeNode {
	if data[len(data)-1] == ',' {
		data = data[:len(data)-1]
	}
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	n := build()
	return n
}

// 331. 验证二叉树的前序序列化
func isValidSerialization(preorder string) bool {
	n := len(preorder)
	stack := []int{1}
	for i := 0; i < n; {
		if len(stack) == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			// 为nil减少栈顶的槽位
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			i++
		} else {
			// 读数字
			for i < n && preorder[i] != ',' {
				i++
			}
			// 为数字减少栈顶的槽位
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			// 在调加一个为2的槽位
			stack = append(stack, 2)
		}
	}
	return len(stack) == 0
}

// 703. 数据流中的第 K 大元素
type KthLargest struct {
	nums []int
	k    int
}

func Constructor3(k int, nums []int) KthLargest {
	kth := KthLargest{nums: make([]int, k), k: k}
	for i := 0; i < k; i++ {
		kth.nums[i] = math.MinInt32
	}
	sort.Ints(nums)
	j := 0
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]
		kth.nums[j] = n
		j++
		if j == k {
			break
		}
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	for i := 0; i < len(this.nums); i++ {
		if val > this.nums[i] {
			this.nums = append(this.nums[:i], append([]int{val}, this.nums[i:]...)...)
			this.nums = this.nums[:this.k]
			break
		}
	}

	return this.nums[this.k-1]
}
