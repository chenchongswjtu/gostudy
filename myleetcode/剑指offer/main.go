package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}

// 剑指 Offer 04. 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	i, j := 0, n-1
	for i >= 0 && i <= m-1 && j >= 0 && j <= n-1 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 剑指 Offer 06. 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	var cur = head
	var ans []int
	for cur != nil {
		ans = append([]int{cur.Val}, ans...)
		cur = cur.Next
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指 Offer 07. 重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= 1000000007
	}

	return dp[n]
}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		sum := (a + b) % 1000000007
		a = b
		b = sum
	}

	return b
}

// 剑指 Offer 11. 旋转数组的最小数字 (二分查找)
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

// 剑指 Offer 12. 矩阵中的路径
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	if len(board[0]) == 0 {
		return false
	}

	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j, index int, visited [][]bool) bool
	dfs = func(i, j, index int, visited [][]bool) bool {
		if index == len(word) {
			return true
		}

		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}

		if visited[i][j] || board[i][j] != word[index] {
			return false
		}

		visited[i][j] = true
		res := dfs(i+1, j, index+1, visited) || dfs(i, j+1, index+1, visited) || dfs(i-1, j, index+1, visited) || dfs(i, j-1, index+1, visited)
		visited[i][j] = false // 回溯
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

// 剑指 Offer 13. 机器人的运动范围
// 广度优先搜索
func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}

	get := func(x int) int {
		res := 0
		for x != 0 {
			res += x % 10
			x = x / 10
		}
		return res
	}

	queue := make([][2]int, 0)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	// 向下和向右的方向数组
	dx := [2]int{0, 1}
	dy := [2]int{1, 0}

	queue = append(queue, [2]int{0, 0})
	visited[0][0] = true
	ans := 1
	for len(queue) > 0 {
		tail := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		x := tail[0]
		y := tail[1]
		for i := 0; i < 2; i++ {
			tx := dx[i] + x
			ty := dy[i] + y
			if tx < 0 || tx >= m || ty < 0 || ty >= n || visited[tx][ty] || get(tx)+get(ty) > k {
				continue
			}
			queue = append(queue, [2]int{tx, ty})
			visited[tx][ty] = true
			ans++
		}
	}

	return ans
}

// 剑指 Offer 14- I. 剪绳子
// 动态规划
func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i < n+1; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[i-j]*j)
		}
	}

	return dp[n]
}

// 剑指 Offer 14- II. 剪绳子
func cuttingRope2(n int) int {
	if n < 4 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res = res * 3 % (1e9 + 7)
		n = n - 3
	}

	return res * n % (1e9 + 7)
}

// 剑指 Offer 25. 合并两个排序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var dummy = &ListNode{}
	var cur = dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			t := &ListNode{Val: l1.Val}
			cur.Next = t
			cur = cur.Next
			l1 = l1.Next
		} else {
			t := &ListNode{Val: l2.Val}
			cur.Next = t
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return helper(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// 包含以A为根的数是否包含B（必须从A开始）
func helper(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}

	if a == nil || a.Val != b.Val {
		return false
	}

	return helper(a.Left, b.Left) && helper(a.Right, b.Right)
}

// 剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := mirrorTree(root.Right)
	r := mirrorTree(root.Left)
	root.Left = l
	root.Right = r
	return root
}

// 剑指 Offer 28. 对称的二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 1 {
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}

		if nums[j]%2 == 0 {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums
}

// 1的个数
func hammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		num &= num - 1
		ans++
	}
	return ans
}

// 剑指 Offer 22. 链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	count := 0
	for fast != nil {
		fast = fast.Next
		count++
		if count == k {
			break
		}
	}

	if count < k {
		return nil
	}

	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

// 剑指 Offer 16. 数值的整数次方
func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(1.0/x, -n)
	}

	if n == 1 {
		return x
	}

	if n == 0 {
		return 1
	}

	if n%2 == 1 {
		t := myPow(x, n/2)
		return t * t * x
	}

	t := myPow(x, n/2)
	return t * t
}

// 剑指 Offer 17. 打印从1到最大的n位数
func printNumbers(n int) []int {
	maxN := func(n int) int {
		res := 1
		for n > 0 {
			res = res * 10
			n = n - 1
		}
		return res - 1
	}

	res := make([]int, 0)
	for i := 1; i <= maxN(n); i++ {
		res = append(res, i)
	}
	return res
}

// 剑指 Offer 24. 反转链表
func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// top k 剑指 Offer 40. 最小的k个数
func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1, k-1)
	return arr[:k]
}
func quickSort(nums []int, left, right int, k int) {
	if left > right {
		return
	}
	i, j, v := left, right, nums[left]
	for i < j {
		for i < j && nums[j] >= v { //如果是求前k大,这里nums[j]>=pivot改成 nums[j]<=pivot
			j--
		}
		for i < j && nums[i] <= v { //如果是求前k大,这里nums[i]<=pivot改成 nums[i]>=pivot即可
			i++
		}
		nums[i], nums[j] = nums[j], nums[i] // 交换
	}
	nums[left], nums[i] = nums[i], nums[left] // 与基准交换
	if k == i {
		return
	}

	if i < k {
		quickSort(nums, i+1, right, k)
	}

	quickSort(nums, left, i-1, k)
}

// 剑指 Offer 42. 连续子数组的最大和
// 定义dp[i]表示数组中前i+1（注意这里的i是从0开始的）个元素构成的连续子数组的最大和。
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	m := dp[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		m = max(m, dp[i])
	}
	return m
}

// 剑指 Offer 31. 栈的压入、弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	i, j := 0, 0
	for i < len(pushed) || j < len(popped) {
		if len(stack) > 0 {
			if j < len(popped) && stack[len(stack)-1] == popped[j] {
				j++
				stack = stack[:len(stack)-1]
				continue
			}
		}

		if i < len(pushed) && j < len(popped) && pushed[i] != popped[j] {
			stack = append(stack, pushed[i])
			i++
		} else {
			i++
			j++
		}
	}

	return len(stack) == 0
}

// 剑指 Offer 39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	count := 0
	var candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// 剑指 Offer 32 - I. 从上到下打印二叉树 (广度优先搜索)
func levelOrder1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	var res []int
	for len(q) > 0 {
		t := q[0]
		res = append(res, t.Val)
		q = q[1:]
		if t.Left != nil {
			q = append(q, t.Left)
		}
		if t.Right != nil {
			q = append(q, t.Right)
		}
	}

	return res
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	var res [][]int
	for len(q) > 0 {
		var t = make([]*TreeNode, len(q))
		copy(t, q)

		q = nil
		var o []int
		for _, node := range t {
			o = append(o, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, o)
	}
	return res
}

// 剑指 Offer 44. 数字序列中某一位的数字
func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}

	digits := 1 // 位数
	begin := 1  // 这个位数最小的数
	count := 9  // 这个位数所以的数的总的位数

	for n > count {
		n -= count
		digits++
		begin *= 10
		count = digits * begin * 9
	}

	num := begin + (n-1)/digits
	nums := strconv.Itoa(num)
	return int(nums[(n-1)%digits] - '0')
}

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	var res [][]int
	for len(q) > 0 {
		var t = make([]*TreeNode, len(q))
		copy(t, q)

		q = nil
		var o []int
		for _, node := range t {
			o = append(o, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, o)
	}

	for i := 0; i < len(res); i++ {
		if i%2 == 1 {
			res[i] = reverse(res[i])
		}
	}
	return res
}

func reverse(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return nums
}

// 剑指 Offer 33. 二叉搜索树的后序遍历序列
func verifyPostorder(postorder []int) bool {
	return verifyPostorderHelper(postorder, 0, len(postorder)-1)
}

func verifyPostorderHelper(postorder []int, l int, r int) bool {
	if l >= r {
		return true
	}

	m := l
	root := postorder[r]

	for postorder[m] < root {
		m++
	}
	t := m

	for t < r {
		if postorder[t] < root {
			return false
		}
		t++
	}

	return verifyPostorderHelper(postorder, l, m-1) && verifyPostorderHelper(postorder, m, r-1)
}

func firstUniqChar(s string) byte {
	m := make(map[byte]int)

	for i := range s {
		m[s[i]]++
	}

	for i := range s {
		if m[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}

// 剑指 Offer 34. 二叉树中和为某一值的路径
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var allPath [][]int
	pathSumHelper(root, target, &allPath, []int{})
	return allPath
}

func pathSumHelper(root *TreeNode, target int, allPath *[][]int, path []int) {
	if root.Left == nil && root.Right == nil {
		if root.Val == target {
			path = append(path, root.Val)
			*allPath = append(*allPath, path)
		}
		return
	}

	t := make([]int, len(path))
	copy(t, path)

	if root.Left != nil {
		path = append(path, root.Val)
		pathSumHelper(root.Left, target-root.Val, allPath, path)
	}

	if root.Right != nil {
		t = append(t, root.Val)
		pathSumHelper(root.Right, target-root.Val, allPath, t)
	}
}

// 剑指 Offer 55 - I. 二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 剑指 Offer 57. 和为s的两个数字
func twoSum(nums []int, target int) []int {
	var ans []int
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		}

		if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return ans
}

// 剑指 Offer 56 - I. 数组中数字出现的次数
func singleNumbers(nums []int) []int {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := set[num]; ok {
			delete(set, num)
			continue
		}
		set[num] = struct{}{}
	}

	var ans []int
	for k := range set {
		ans = append(ans, k)
	}
	return ans
}

// 剑指 Offer 45. 把数组排成最小的数
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		si := strconv.Itoa(nums[i])
		sj := strconv.Itoa(nums[j])
		return si+sj < sj+si
	})

	var res string
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return res
}

// 剑指 Offer 57 - II. 和为s的连续正数序列
// ***
func findContinuousSequence(target int) [][]int {
	var res [][]int
	for l, r := 1, 2; l < r; {
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			var one []int
			for i := l; i <= r; i++ {
				one = append(one, i)
			}
			res = append(res, one)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return res
}

// 剑指 Offer 46. 把数字翻译成字符串
func translateNum(num int) int {
	var count int
	translateNumHelper(strconv.Itoa(num), 0, &count)
	return count
}

func translateNumHelper(numStr string, start int, count *int) {
	if start >= len(numStr) {
		*count++
		return
	}

	if numStr[start] == '0' {
		translateNumHelper(numStr, start+1, count)
		return
	}

	for i := start; i < len(numStr); i++ {
		s := numStr[start : i+1]
		n, _ := strconv.Atoi(s)
		if n > 25 {
			break
		}
		translateNumHelper(numStr, i+1, count)
	}
}

// 剑指 Offer 52. 两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ma := make(map[*ListNode]*ListNode)
	mb := make(map[*ListNode]*ListNode)

	cura := headA
	for cura != nil {
		ma[cura] = cura.Next
		cura = cura.Next
	}

	curb := headB
	for curb != nil {
		mb[curb] = curb.Next
		curb = curb.Next
	}

	cura = headA
	for cura != nil {
		next, ok := mb[cura]
		if ok {
			if ma[cura] == next {
				return cura
			}
		}
		cura = cura.Next
	}

	return nil
}

// 剑指 Offer 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	ss := strings.Split(s, " ")

	var ss1 []string
	for i := 0; i < len(ss); i++ {
		if ss[i] != "" {
			ss1 = append(ss1, ss[i])
		}
	}

	for i, j := 0, len(ss1)-1; i < j; {
		ss1[i], ss1[j] = ss1[j], ss1[i]
		i++
		j--
	}
	return strings.Join(ss1, " ")
}

// 剑指 Offer 47. 礼物的最大价值
func maxValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else {
			//等与target时候也-1，收缩右边界来锁定左侧边界
			r = m - 1
		}
	}
	left := l

	l, r = 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		} else {
			//等与target时候也+1，收缩左边界来锁定右侧边界
			l = m + 1
		}
	}
	right := r

	return right - left + 1
}

// 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		m := i + (j-i)/2
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

// 剑指 Offer 54. 二叉搜索树的第k大节点
func kthLargest(root *TreeNode, k int) int {
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
	return res[len(res)-k]
}

// 剑指 Offer 48. 最长不含重复字符的子字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	l, r := 0, 0
	maxLen := 0
	set := make(map[byte]struct{})

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for l < len(s) {
		for r < len(s) {
			if _, ok := set[s[r]]; ok {
				break
			}
			set[s[r]] = struct{}{}
			r++
		}
		maxLen = max(maxLen, r-l)

		// 删除前面重复的
		delete(set, s[l])
		l++
	}
	return maxLen
}

// 剑指 Offer 49. 丑数
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 剑指 Offer 66. 构建乘积数组
func constructArr(a []int) []int {
	left := make([]int, len(a))
	right := make([]int, len(a))

	for i := range a {
		if i == 0 {
			left[i] = 1
		} else {
			left[i] = a[i-1] * left[i-1]
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		if i == len(a)-1 {
			right[i] = 1
		} else {
			right[i] = a[i+1] * right[i+1]
		}
	}

	ans := make([]int, len(a))
	for i := range a {
		ans[i] = left[i] * right[i]
	}
	return ans
}

// 剑指 Offer 60. n个骰子的点数
func dicesProbability(n int) []float64 {
	// 初始是1个骰子情况下的点数之和情况，就只有6个结果，所以用dp的初始化的size是6个
	dp := make([]float64, 6)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1.0 / 6.0
	}

	//从第2个骰子开始，这里n表示n个骰子，先从第二个的情况算起，然后再逐步求3个、4个···n个的情况
	//i表示当总共i个骰子时的结果
	for i := 2; i <= n; i++ {
		//每次的点数之和范围会有点变化，点数之和的值最大是i*6，最小是i*1，i之前的结果值是不会出现的；
		//比如i=3个骰子时，最小就是3了，不可能是2和1，所以点数之和的值的个数是6*i-(i-1)，化简：5*i+1
		//当有i个骰子时的点数之和的值数组先假定是t
		t := make([]float64, 5*i+1)
		//从i-1个骰子的点数之和的值数组入手，计算i个骰子的点数之和数组的值
		//先拿i-1个骰子的点数之和数组的第j个值，它所影响的是i个骰子时的t[j+k]的值
		for j := 0; j < len(dp); j++ {
			//比如只有1个骰子时，dp[1]是代表当骰子点数之和为2时的概率，它会对当有2个骰子时的点数之和为3、4、5、6、7、8产生影响，
			//因为当有一个骰子的值为2时，另一个骰子的值可以为1~6，产生的点数之和相应的就是3~8；
			//比如dp[2]代表点数之和为3，它会对有2个骰子时的点数之和为4、5、6、7、8、9产生影响；
			//所以k在这里就是对应着第i个骰子出现时可能出现六种情况，这里可能画一个K神那样的动态规划逆推的图就好理解很多
			for k := 0; k < 6; k++ {
				//这里记得是加上dp数组值与1/6的乘积，1/6是第i个骰子投出某个值的概率
				t[j+k] += dp[j] * (1.0 / 6.0)
			}
		}
		//i个骰子的点数之和全都算出来后，要将temp数组移交给dp数组，dp数组就会代表i个骰子时的可能出现的点数之和的概率；用于计算i+1个骰子时的点数之和的概率
		dp = t
	}

	return dp
}

// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if (l-r > 1) || (l-r < -1) {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 剑指 Offer 61. 扑克牌中的顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)

	count := 0
	first := true
	for i, num := range nums {
		if num == 0 {
			count++
			continue
		}

		if num != 0 && first {
			first = false
			continue
		}

		diff := nums[i] - nums[i-1]
		if diff == 0 {
			return false
		}
		if diff == 1 {
			continue
		}

		if count <= 0 {
			return false
		}

		count = count - (diff - 1)
	}

	return count >= 0
}

// 剑指 Offer 63. 股票的最大利润
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := 0
	m := prices[0]

	for i := 1; i < len(prices); i++ {
		m = min(m, prices[i-1])
		res = max(res, prices[i]-m)
	}

	return res
}

// 剑指 Offer 62. 圆圈中最后剩下的数字
func lastRemaining(n int, m int) int {
	return f(n, m)
}

func f(n int, m int) int {
	if n == 1 {
		return 0
	}
	x := f(n-1, m)
	return (x + m) % n
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//剑指 Offer 35. 复杂链表的复制
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	cur := head

	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}

	return m[head]
}

// 剑指 Offer 59 - I. 滑动窗口的最大值
// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	var q []int
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] { // 单调递减
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	var ans []int
	ans = append(ans, nums[q[0]])

	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k { // 最大的大于K
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

// 最小栈
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	top := s.minStack[len(s.minStack)-1]
	s.minStack = append(s.minStack, min(x, top))
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) Min() int {
	return s.minStack[len(s.minStack)-1]
}

// 顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 顺时针方向
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

// 剑指 Offer 59 - II. 队列的最大值
type MaxQueue struct {
	q   []int
	max []int
}

func Constructor1() MaxQueue {
	return MaxQueue{
		q:   make([]int, 0),
		max: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	for len(this.max) != 0 && value > this.max[len(this.max)-1] { // 递减
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.q) != 0 {
		n := this.q[0]
		this.q = this.q[1:]
		if this.max[0] == n {
			this.max = this.max[1:]
		}
	}
	return n
}

// 表示数值的字符串
type State int
type CharType int

const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		STATE_INITIAL: map[CharType]State{
			CHAR_SPACE:  STATE_INITIAL,
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
			CHAR_SIGN:   STATE_INT_SIGN,
		},
		STATE_INT_SIGN: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
		},
		STATE_INTEGER: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_EXP:    STATE_EXP,
			CHAR_POINT:  STATE_POINT,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT_WITHOUT_INT: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_EXP: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SIGN:   STATE_EXP_SIGN,
		},
		STATE_EXP_SIGN: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
		STATE_EXP_NUMBER: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SPACE:  STATE_END,
		},
		STATE_END: map[CharType]State{
			CHAR_SPACE: STATE_END,
		},
	}
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}

// 数组中的逆序对 (归并排序)
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start int, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	count := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	var tmp []int
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i]) // nums[i] > nums[mid+1:j-1] 就有j-（mid+1）个逆序对
			count += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		count += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return count
}
