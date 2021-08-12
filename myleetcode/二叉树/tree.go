package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var null = NULL

func main() {
	fmt.Println(balanceBST(Ints2TreeNode([]int{1, null, 2, null, 3, null, 4, null, null})))
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
		return targetSum == root.Val
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
	pre := preOrder(root)
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

// 包含a,b根节点验证是否相同
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

func (kth *KthLargest) Add(val int) int {
	for i := 0; i < len(kth.nums); i++ {
		if val > kth.nums[i] {
			kth.nums = append(kth.nums[:i], append([]int{val}, kth.nums[i:]...)...)
			kth.nums = kth.nums[:kth.k]
			break
		}
	}

	return kth.nums[kth.k-1]
}

// 863. 二叉树中所有距离为 K 的结点
func distanceK(root, target *TreeNode, k int) (ans []int) {
	// 从root出发dfs，记录每个节点的父节点
	parents := map[int]*TreeNode{}
	var findParents func(node *TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}

	findParents(root)

	// 从target出发dfs，寻找所有深度为k的节点
	// from 为来源节点
	var findAns func(node *TreeNode, from *TreeNode, depth int)
	findAns = func(node *TreeNode, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k {
			ans = append(ans, node.Val)
		}
		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}
		if parents[node.Val] != from {
			findAns(parents[node.Val], node, depth+1)
		}
	}

	findAns(target, nil, 0)
	return
}

// 865. 具有所有最深节点的最小子树
// 思路：从每个树开始，获得当前节点的左右子树的最大深度
// 深度相同，说明最深的节点在这个节点两边，那这个节点就是结果
// 如果深度不相同，则去深度大的子树继续判断，最终就能得到结果
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// 获取当前节点的左右子树的最大深度
	leftMaxDepth := getHeight(root.Left)
	rightMaxDepth := getHeight(root.Right)

	// 如果两边最大深度相同，则这个节点就是结果
	if leftMaxDepth == rightMaxDepth {
		return root
	}

	// 不相等，那就去深度大的子树那边继续找
	if leftMaxDepth > rightMaxDepth {
		return subtreeWithAllDeepest(root.Left)
	}

	return subtreeWithAllDeepest(root.Right)
}

// 897. 递增顺序搜索树
func increasingBST(root *TreeNode) *TreeNode {
	var nums []int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node != nil {
			inOrder(node.Left)
			nums = append(nums, node.Val)
			inOrder(node.Right)
		}
	}

	inOrder(root)

	dummy := &TreeNode{}
	cur := dummy
	for _, num := range nums {
		cur.Right = &TreeNode{Val: num}
		cur = cur.Right
	}

	return dummy.Right
}

// 951. 翻转等价二叉树
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == root2 {
		return true
	}

	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}

	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}

//958. 二叉树的完全性检验
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q = []*TreeNode{root}
	var isFin bool
	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		if t.Left != nil {
			if isFin { // 如果已经isFin为true再遇到不为nil的节点直接返回false
				return false
			}
			q = append(q, t.Left)
		} else { // 遇到第一个nil的设置isFin为true
			isFin = true
		}

		if t.Right != nil {
			if isFin {
				return false
			}
			q = append(q, t.Right)
		} else {
			isFin = true
		}
	}
	return true
}

// 971. 翻转二叉树以匹配先序遍历
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var res []int
	// 使用res记录需要反转的节点的值

	var dfs func(root *TreeNode, voyage *[]int) bool

	dfs = func(root *TreeNode, voyage *[]int) bool {
		if root == nil {
			return true
		}
		// 根节点的值和voyage第0个元素不符，无法通过反转节点使使先序遍历符合voyage
		if root.Val != (*voyage)[0] {
			res = []int{-1}
			return false
		}
		// 当且仅当根节点的左右节点都不为空，且右儿子的值等于voyage的第1个元素，需要反转根节点的左右儿子
		if root.Left != nil && root.Right != nil && root.Right.Val == (*voyage)[1] {
			res = append(res, root.Val)
			root.Left, root.Right = root.Right, root.Left
		}
		// 消耗掉voyage第0个元素
		*voyage = (*voyage)[1:]
		// 如果对左儿子递归的结果为false，就进行剪枝操作
		ok := dfs(root.Left, voyage)
		if !ok {
			return false
		}
		return dfs(root.Right, voyage)
	}

	dfs(root, &voyage)
	return res
}

// 993. 二叉树的堂兄弟节点
func isCousins(root *TreeNode, x, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	var dfs func(node *TreeNode, parent *TreeNode, depth int)
	dfs = func(node *TreeNode, parent *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}

		if xFound && yFound {
			return
		}

		dfs(node.Left, node, depth+1)
		if xFound && yFound {
			return
		}
		dfs(node.Right, node, depth+1)
	}

	dfs(root, nil, 0)

	return xDepth == yDepth && xParent != yParent
}

// 1008. 前序遍历构造二叉搜索树
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: preorder[0]}
	var i int
	for i = 1; i < len(preorder); i++ {
		if preorder[i] >= preorder[0] {
			break
		}
	}

	node.Left = bstFromPreorder(preorder[1:i])
	node.Right = bstFromPreorder(preorder[i:])
	return node
}

// 1022. 从根到叶的二进制数之和
func sumRootToLeaf(root *TreeNode) int {
	var sum int
	var dfs func(node *TreeNode, prev int)
	// prev为这个节点前面数的值
	dfs = func(node *TreeNode, prev int) {
		// 根节点计算数据到sum
		if node.Left == nil && node.Right == nil {
			sum += prev*2 + node.Val
			return
		}

		if node.Left != nil {
			dfs(node.Left, prev*2+node.Val)
		}

		if node.Right != nil {
			dfs(node.Right, prev*2+node.Val)
		}
	}

	dfs(root, 0)
	return sum
}

// 1026. 节点与其祖先之间的最大差值
func maxAncestorDiff(root *TreeNode) int {
	var maxDiff int
	var dfs func(node *TreeNode, min, max int)

	dfs = func(node *TreeNode, min, max int) {
		if node == nil {
			return
		}
		val := node.Val
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
		if max-min > maxDiff {
			maxDiff = max - min
		}
		if node.Left != nil {
			dfs(node.Left, min, max)
		}
		if node.Right != nil {
			dfs(node.Right, min, max)
		}
	}

	dfs(root, root.Val, root.Val)
	return maxDiff
}

// 1028. 从先序遍历还原二叉树
func recoverFromPreorder(traversal string) *TreeNode {
	var path []*TreeNode
	var pos int
	for pos < len(traversal) {
		level := 0
		for traversal[pos] == '-' {
			level++
			pos++
		}
		value := 0
		// 获得值
		for ; pos < len(traversal) && traversal[pos] >= '0' && traversal[pos] <= '9'; pos++ {
			value = value*10 + int(traversal[pos]-'0')
		}
		node := &TreeNode{Val: value}
		if level == len(path) {
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			path = path[:level]
			path[len(path)-1].Right = node
		}
		path = append(path, node)
	}

	if len(path) > 0 {
		return path[0]
	}
	return nil
}

// 998. 最大二叉树 II
// 如果val > root.Val, 直接使之成为根，反之，插入到右子树中
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		return &TreeNode{Val: val, Left: root, Right: nil}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

// 面试题 04.02. 最小高度树
//func sortedArrayToBST(nums []int) *TreeNode {
//	if len(nums) == 0 {
//		return nil
//	}
//
//	l, r := 0, len(nums)-1
//	m := (l + r) / 2
//	root := &TreeNode{Val: nums[m]}
//	root.Left = sortedArrayToBST(nums[:m])
//	root.Right = sortedArrayToBST(nums[m+1:])
//	return root
//}

// 1038. 把二叉搜索树转换为累加树
func bstToGst(root *TreeNode) *TreeNode {
	var nodes []*TreeNode
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		nodes = append(nodes, node)
		inOrder(node.Right)
	}

	inOrder(root)

	sum := 0
	for i := len(nodes) - 1; i >= 0; i-- {
		sum += nodes[i].Val
		nodes[i].Val = sum
	}

	return root
}

// 1161. 最大层内元素和
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q = []*TreeNode{root}
	var maxValue = math.MinInt64
	var maxLevel = 1
	var level = 0
	for len(q) > 0 {
		t := q
		q = nil
		sum := 0
		level++
		for _, node := range t {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > maxValue {
			maxValue = sum
			maxLevel = level
		}
	}
	return maxLevel
}

// 1080. 根到叶路径上的不足节点
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var l, r *TreeNode

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		if root.Val < limit {
			return nil
		}
		return root
	}

	l = sufficientSubset(root.Left, limit-root.Val)
	r = sufficientSubset(root.Right, limit-root.Val)
	// 左右子树删除之后都为nil，这时root也需要判断是否要删，又由于之前已经-root.val，
	// 所以root一定要删，返回nil
	if l == nil && r == nil {
		return nil
	}

	root.Left = l
	root.Right = r
	return root
}

// 1104. 二叉树寻路
// 2^(i-1)+2^i-1-label
// i 为行数
// 偶数行将label转化为反方向
func getReverse(label, row int) int {
	return 1<<(row-1) + 1<<row - 1 - label
}

func pathInZigZagTree(label int) (path []int) {
	row, rowStart := 1, 1
	for rowStart*2 <= label {
		row++
		rowStart *= 2
	}
	// 根据最后行是偶数将其转化为正常的label（全部是从左到右）
	if row%2 == 0 {
		label = getReverse(label, row)
	}
	for row > 0 {
		if row%2 == 0 {
			path = append(path, getReverse(label, row))
		} else {
			path = append(path, label)
		}
		row--
		label >>= 1
	}
	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}
	return
}

// 1123. 最深叶节点的最近公共祖先
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var maxDepth int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(height(node.Left), height(node.Right)) + 1
	}

	var contain func(node *TreeNode, depth int) bool
	contain = func(node *TreeNode, depth int) bool {
		if node == nil {
			return false
		}
		// 叶子节点
		if node.Left == nil && node.Right == nil {
			// 检查是否是最深的叶子节点
			return depth == maxDepth
		}
		return contain(node.Left, depth+1) || contain(node.Right, depth+1)
	}

	if root == nil {
		return nil
	}

	maxDepth = height(root)

	// 左右子树都包含最深叶子节点
	if contain(root.Left, 2) && contain(root.Right, 2) {
		return root
	} else if contain(root.Left, 2) { // 只有左子树包含，递归寻找最近的祖先
		return lcaDeepestLeaves(root.Left)
	} else if contain(root.Right, 2) {
		return lcaDeepestLeaves(root.Right)
	} else { // 左右子树都不包含，直接返回根节点
		return root
	}
}

// 1302. 层数最深叶子节点的和
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var maxDepth int
	var sum int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 获得树的层数
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(getDepth(node.Left), getDepth(node.Right)) + 1
	}

	maxDepth = getDepth(root)
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			if depth == maxDepth {
				sum += node.Val
			}
			return
		}

		if node.Left != nil {
			dfs(node.Left, depth+1)
		}

		if node.Right != nil {
			dfs(node.Right, depth+1)
		}
	}

	// 根节点为第一层
	dfs(root, 1)
	return sum
}

// 1382. 将二叉搜索树变平衡
func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var nums []int
	var inOrder func(node *TreeNode)
	// 中序遍历将二叉搜索树按照从小到大的顺序遍历
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		nums = append(nums, node.Val)
		inOrder(node.Right)
	}

	// 使用二分法来创建平衡树
	var build func(nums []int) *TreeNode
	build = func(nums []int) *TreeNode {
		l, r := 0, len(nums)-1
		if l > r {
			return nil
		}
		m := (l + r) / 2
		return &TreeNode{Val: nums[m], Left: build(nums[:m]), Right: build(nums[m+1:])}
	}

	inOrder(root)

	return build(nums)
}

// 894. 所有可能的满二叉树
// 代码里d[n]为节点数是n的二叉树数组，i则该二叉树代表左侧分配的节点数，n - i - 1则为右侧分配的节点数，
// 再遍历组合左右节点数对应的二叉树数组即可得到n个节点时的二叉树数组。
func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return []*TreeNode{}
	}
	d := [20][]*TreeNode{1: {&TreeNode{}}}
	for n := 3; n <= N; n += 2 {
		for i := 1; i < n; i += 2 {
			for _, left := range d[i] {
				for _, right := range d[n-i-1] {
					d[n] = append(d[n], &TreeNode{Left: left, Right: right})
				}
			}
		}
	}
	return d[N]
}

// 988. 从叶结点开始的最小字符串
func smallestFromLeaf(root *TreeNode) string {
	ans := "~" // 比字符串都大
	var dfs func(node *TreeNode, s string)
	dfs = func(node *TreeNode, s string) {
		if node == nil {
			return
		}
		s += string(rune('a' + node.Val))
		if node.Left == nil && node.Right == nil {
			s = reverseString(s)             // 翻转字符串
			if strings.Compare(s, ans) < 0 { // 比较大小
				ans = s
			}
		}
		dfs(node.Left, s)
		dfs(node.Right, s)
	}

	dfs(root, "")
	return ans
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// 1361. 验证二叉树
// 先找根节点（入度为0）
// 再bfs遍历，看是否遍历完
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	var in = make([]int, n) // 入度数组
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			in[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			in[rightChild[i]]++
		}
	}

	root := -1
	for i := 0; i < n; i++ {
		if in[i] == 0 { // 入度为0的为根节点
			root = i
			break
		}
	}

	if root == -1 { // 没有入度为0的节点
		return false
	}

	var seen = make(map[int]struct{}) // 已经遍历到
	var queue = []int{root}
	seen[root] = struct{}{}

	// bfs 宽度优先遍历
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		if leftChild[e] != -1 {
			if _, ok := seen[leftChild[e]]; ok { // 已经遍历到过
				return false
			}
			seen[leftChild[e]] = struct{}{}
			queue = append(queue, leftChild[e])
		}

		if rightChild[e] != -1 {
			if _, ok := seen[rightChild[e]]; ok { // 已经遍历到过
				return false
			}
			seen[rightChild[e]] = struct{}{}
			queue = append(queue, rightChild[e])
		}
	}

	// 是否遍历完
	return len(seen) == n
}

// 1325. 删除给定值的叶子节点
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	// 删除左右子树之后，左右子树为nil，继续看根节点与是否可以删除
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}

// 1339. 分裂二叉树的最大乘积
func maxProduct(root *TreeNode) int {
	var sum int
	var half int

	// 计算这个树上网总值
	var dfs1 func(node *TreeNode)
	dfs1 = func(node *TreeNode) {
		if node == nil {
			return
		}

		sum += node.Val
		dfs1(node.Left)
		dfs1(node.Right)
	}

	// 计算这个子树的总数
	var dfs2 func(node *TreeNode) int
	dfs2 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		cur := dfs2(node.Left) + dfs2(node.Right) + node.Val
		if math.Abs(float64(cur*2-sum)) < math.Abs(float64(half*2-sum)) { // 计算最接近sum一半的值
			half = cur
		}
		return cur
	}

	dfs1(root)
	dfs2(root)
	return half * (sum - half) % 1000000007
}

// 1372. 二叉树中的最长交错路径
func longestZigZag(root *TreeNode) int {
	var maxLen int
	const left, right = 0, 1

	// dir 下一步的方向
	var dfs func(node *TreeNode, dir int, length int)
	dfs = func(node *TreeNode, dir int, length int) {
		maxLen = max(maxLen, length)

		if dir == left {
			if node.Left != nil {
				dfs(node.Left, right, length+1)
			}
			if node.Right != nil { // 不按照下一步的方向，重置l为1
				dfs(node.Right, left, 1)
			}
		} else {
			if node.Right != nil {
				dfs(node.Right, left, length+1)
			}
			if node.Left != nil {
				dfs(node.Left, right, 1)
			}
		}
	}

	if root == nil {
		return 0
	}

	dfs(root, left, 0)
	dfs(root, right, 0)

	return maxLen
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1367. 二叉树中的列表
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	// 以head和root为起点进行遍历
	var dfs func(*ListNode, *TreeNode) bool
	dfs = func(head *ListNode, root *TreeNode) bool {
		// 链表已经全部匹配完，匹配成功
		if head == nil {
			return true
		}
		// 二叉树访问到了空节点，匹配失败
		if root == nil {
			return false
		}
		// 当前匹配的二叉树上节点的值与链表节点的值不相等，匹配失败
		if head.Val != root.Val {
			return false
		}
		return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
	}

	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
