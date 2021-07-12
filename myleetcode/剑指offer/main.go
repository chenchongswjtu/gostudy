package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reverseWords("a good   example"))
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
