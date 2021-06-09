package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calcEquation1([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
}

// 3. 无重复字符的最长子串(滑动窗口)
func lengthOfLongestSubstring(s string) int {
	maxLen, curLen := 0, 0
	set := make(map[uint8]struct{})
	var left = 0
	for i := 0; i < len(s); i++ {
		_, ok := set[s[i]]
		for ok {
			delete(set, s[left])
			left++
			curLen--
			_, ok = set[s[i]]
		}

		set[s[i]] = struct{}{}
		curLen++
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 {
		if n2%2 == 0 {
			return float64(nums2[(n2-1)/2]+nums2[(n2-1)/2+1]) / 2
		} else {
			return float64(nums2[n2/2])
		}
	}

	if n2 == 0 {
		if n1%2 == 0 {
			return float64(nums1[(n1-1)/2]+nums1[(n1-1)/2+1]) / 2
		} else {
			return float64(nums1[n1/2])
		}
	}

	if (n1+n2)%2 == 0 {
		target1, target2 := (n1+n2-1)/2+1, (n1+n2-1)/2+2
		t1, t2 := -1, -1
		count := 0
		i, j := 0, 0
		ans := 0
		for i < n1 || j < n2 {
			if i < n1 && j < n2 {
				if nums1[i] <= nums2[j] {
					ans = nums1[i]
					i++
					count++
				} else {
					ans = nums2[j]
					j++
					count++
				}
			} else if i < n1 {
				ans = nums1[i]
				i++
				count++
			} else if j < n2 {
				ans = nums2[j]
				j++
				count++
			}

			if count == target1 {
				t1 = ans
			}

			if count == target2 {
				t2 = ans
			}

			if t1 != -1 && t2 != -1 {
				return float64(t1+t2) / 2
			}
		}
	} else {
		target := (n1+n2)/2 + 1
		count := 0
		i, j := 0, 0
		ans := 0
		for i < n1 || j < n2 {
			if i < n1 && j < n2 {
				if nums1[i] <= nums2[j] {
					ans = nums1[i]
					i++
					count++
				} else {
					ans = nums2[j]
					j++
					count++
				}
			} else if i < n1 {
				ans = nums1[i]
				i++
				count++
			} else if j < n2 {
				ans = nums2[j]
				j++
				count++
			}
			if count == target {
				return float64(ans)
			}
		}
	}
	return -1
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if (n1+n2)&1 == 0 {
		l := find(nums1, 0, nums2, 0, (n1+n2)/2)
		r := find(nums1, 0, nums2, 0, (n1+n2)/2+1)
		return float64(l+r) / 2
	}

	return float64(find(nums1, 0, nums2, 0, (n1+n2)/2+1))
}

func find(nums1 []int, i int, nums2 []int, j int, k int) int {
	if len(nums1)-i > len(nums2)-j {
		return find(nums2, j, nums1, i, k)
	}

	if len(nums1) == i {
		return nums2[j+k-1]
	}

	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	si := min(len(nums1), i+k/2)
	sj := j + k - k/2
	if nums1[si-1] > nums2[sj-1] {
		return find(nums1, i, nums2, sj, k-(sj-j))
	} else {
		return find(nums1, si, nums2, j, k-(si-i))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 11. 盛最多水的容器
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		a := min(height[l], height[r]) * (r - l)
		ans = max(ans, a)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

// 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)

	// 枚举a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for n1 := 0; n1 < n; n1++ {
		if n1 > 0 && nums[n1] == nums[n1-1] {
			continue
		}
		n3 := n - 1
		target := -nums[n1]
		for n2 := n1 + 1; n2 < n; n2++ {
			if n2 > n1+1 && nums[n2] == nums[n2-1] {
				continue
			}
			for n2 < n3 && nums[n2]+nums[n3] > target {
				n3--
			}
			if n2 == n3 {
				break
			}
			if nums[n2]+nums[n3] == target {
				ans = append(ans, []int{nums[n1], nums[n2], nums[n3]})
			}
		}
	}
	return ans
}

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	letter := make(map[string][]string)
	letter["2"] = []string{"a", "b", "c"}
	letter["3"] = []string{"d", "e", "f"}
	letter["4"] = []string{"g", "h", "i"}
	letter["5"] = []string{"j", "k", "l"}
	letter["6"] = []string{"m", "n", "o"}
	letter["7"] = []string{"p", "q", "r", "s"}
	letter["8"] = []string{"t", "u", "v"}
	letter["9"] = []string{"w", "x", "y", "z"}

	var ans []string
	letterCombinationsHelper(letter, digits, 0, "", &ans)
	return ans
}

func letterCombinationsHelper(letter map[string][]string, digits string, i int, s string, ans *[]string) {
	if i == len(digits) {
		*ans = append(*ans, s)
		return
	}

	for _, v := range letter[digits[i:i+1]] {
		t := s
		s += v
		letterCombinationsHelper(letter, digits, i+1, s, ans)
		s = t
	}
}

// 20. 有效的括号
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}

	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		}

		if c == ')' || c == ']' || c == '}' {
			if len(stack) == 0 || pairs[c] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isValid1(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for n > 0 {
		// n大于链表长度
		if fast == nil {
			return nil
		}
		fast = fast.Next
		n--
	}
	// n等于链表长度，为删除链表第一个节点
	if fast == nil {
		return head.Next
	}

	// fast.Next为nil，slow到删除的节点的前一个
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}

// 21. 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	var cur = head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{Val: l1.Val, Next: nil}
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val, Next: nil}
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		cur.Next = l2
	}

	if l2 == nil {
		cur.Next = l1
	}

	return head.Next
}

// 31. 下一个排列
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 不是最后一个排列
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	x, y := j, len(nums)-1
	for x < y {
		nums[x], nums[y] = nums[y], nums[x]
		x++
		y--
	}
}

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	x, y := -1, -1
	for i <= j {
		if x == -1 && nums[i] == target {
			x = i
		}

		if y == -1 && nums[j] == target {
			y = j
		}

		if x != -1 && y != -1 {
			break
		}

		if x == -1 {
			i++
		}
		if y == -1 {
			j--
		}
	}

	return []int{x, y}
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange1(nums []int, target int) []int {
	x, y := -1, -1
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			if m == 0 || m-1 >= 0 && nums[m-1] != target {
				x = m
			}
		}

		if target <= nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	l, r = 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			if m == len(nums)-1 || m+1 < len(nums) && nums[m+1] != target {
				y = m
			}
		}

		if nums[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return []int{x, y}
}

// 23. 合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	var mergeTwoLists func(l1 *ListNode, l2 *ListNode) *ListNode
	mergeTwoLists = func(l1 *ListNode, l2 *ListNode) *ListNode {
		head := &ListNode{}
		cur := head
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				node := &ListNode{Val: l1.Val, Next: nil}
				cur.Next = node
				cur = cur.Next
				l1 = l1.Next
			} else {
				node := &ListNode{Val: l2.Val, Next: nil}
				cur.Next = node
				cur = cur.Next
				l2 = l2.Next
			}
		}

		if l1 == nil {
			cur.Next = l2
		}

		if l2 == nil {
			cur.Next = l1
		}

		return head.Next
	}

	var merge func(lists []*ListNode, l, r int) *ListNode
	merge = func(lists []*ListNode, l, r int) *ListNode {
		if l == r {
			return lists[l]
		}

		if l > r {
			return nil
		}

		m := (l + r) >> 1
		return mergeTwoLists(merge(lists, l, m), merge(lists, m+1, r))
	}

	return merge(lists, 0, len(lists)-1)
}

// 48. 旋转图像(顺时针旋转90度)
// m[x][y] = m[y][n-1-x] (旋转90)
// 先上下调整  m[x][y] = m[n-1-x][y]
// 在对角线调整m[n-1-x][y] =m[y][n-1-x]
// 则为 m[x][y] = m[y][n-1-x] (旋转90)
func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}

	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 49. 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}

	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

// 55. 跳跃游戏
func canJump(nums []int) bool {
	n := len(nums)
	max := 0
	for i := 0; i < n; i++ {
		if i > max {
			break
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}

		if max >= n-1 {
			return true
		}
	}

	return false
}

// 75. 颜色分类
func sortColors(nums []int) {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var ans []int
	for _, color := range []int{0, 1, 2} {
		if c, ok := m[color]; ok {
			for i := 0; i < c; i++ {
				ans = append(ans, color)
			}
		}
	}

	copy(nums, ans)
}

// 75. 颜色分类(原地排序)
func sortColors1(nums []int) {
	n := len(nums)
	i0, i2 := 0, n-1

	for i := 0; i <= i2; {
		if nums[i] == 0 && i != i0 {
			nums[i], nums[i0] = nums[i0], nums[i]
			i0++
		} else if nums[i] == 2 && i != i2 {
			nums[i], nums[i2] = nums[i2], nums[i]
			i2--
		} else {
			i++
		}
	}
}

// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	maxLen := 1
	curLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			curLen++
			if curLen > maxLen {
				maxLen = curLen
			}
		} else if nums[i]-nums[i-1] == 0 {
			// 相等，不执行操作，只是i++
		} else {
			curLen = 1
		}
	}
	return maxLen
}

// 128. 最长连续序(o(n)时间复杂度)
func longestConsecutive1(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			curNum := num
			cur := 1
			for numSet[curNum+1] {
				curNum++
				cur++
			}
			if longest < cur {
				longest = cur
			}
		}
	}
	return longest
}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	var hasCycle func(head *ListNode) bool
	hasCycle = func(head *ListNode) bool {
		if head == nil || head.Next == nil {
			return false
		}
		fast, slow := head, head
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				return true
			}
		}
		return false
	}

	if !hasCycle(head) {
		return nil
	}

	visited := make(map[*ListNode]bool)
	cur := head
	for {
		visited[cur] = true
		next := cur.Next
		if visited[next] {
			return next
		}
		cur = next
	}
}

// LRUCache 146. LRU 缓存机制
type LRUCache struct {
	cache map[int]int
	order []int // 记录key的顺序
	cap   int   // 记录容量
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]int),
		order: make([]int, 0),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// 判断是否是在缓存中
	if v, ok := this.cache[key]; ok {
		pos := 0
		// 在order中找到这个key
		for i, k := range this.order {
			if k == key {
				pos = i
			}
		}
		// 将key的顺序更新到最新
		this.order = append(this.order[:pos], this.order[pos+1:]...)
		this.order = append(this.order, key)
		return v
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 先判断缓存中是否已经存在
	if _, ok := this.cache[key]; ok {
		pos := 0
		for i, k := range this.order {
			if k == key {
				pos = i
			}
		}
		// 存在更新order的顺序，并且更新这个key的值
		this.order = append(this.order[:pos], this.order[pos+1:]...)
		this.order = append(this.order, key)
		this.cache[key] = value
		return
	}

	// 判断缓存的数量是否小于容量
	if len(this.cache) < this.cap {
		// 小于直接添加
		this.cache[key] = value
		this.order = append(this.order, key)
		return
	}

	// 等于直接删除第一个顺序中key的缓存，将新的key添加到orderer的最后存入缓存
	delete(this.cache, this.order[0])
	this.order = append(this.order[1:], key)
	this.cache[key] = value
}

// 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	m := make(map[*ListNode]struct{})

	for curA != nil {
		m[curA] = struct{}{}
		curA = curA.Next
	}

	for curB != nil {
		if _, ok := m[curB]; ok {
			return curB
		}
		curB = curB.Next
	}

	return nil
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur = head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 200. 岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	x, y := len(grid), len(grid[0])
	ans := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= x || j < 0 || j >= y || grid[i][j] == '0' {
			return
		}

		// 遍历到的1设置为0
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}

	return ans
}

// 240. 搜索二维矩阵 II (从矩阵的右上角开始)
func searchMatrix(matrix [][]int, target int) bool {
	x := len(matrix)
	if x == 0 {
		return false
	}

	y := len(matrix[0])
	if y == 0 {
		return false
	}

	i := 0
	j := y - 1

	for i >= 0 && i < x && j >= 0 && j < y {
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

// 279. 完全平方数 (动态规划)
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0

	maxSqrtIndex := int(math.Sqrt(float64(n))) + 1
	squNums := make([]int, maxSqrtIndex)
	for i := 1; i < maxSqrtIndex; i++ {
		squNums[i] = i * i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j < maxSqrtIndex; j++ {
			if i < squNums[j] {
				break
			}
			dp[i] = min(dp[i], dp[i-squNums[j]]+1)
		}
	}
	return dp[n]
}

// 322. 零钱兑换
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			if dp[i-coin] == math.MaxInt64 {
				continue
			}

			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

// 309. 最佳买卖股票时机含冷冻期
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	// dp[i][0]: 手上持有股票的最大收益
	// dp[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// dp[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i]) // 不做，上一次不在冷冻期买股票
		dp[i][1] = dp[i-1][0] + prices[i]                // 前一天有股票再买
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])           // 前一天不在冷冻期不做，前一天是在冷冻期

	}

	return max(dp[n-1][1], dp[n-1][2])
}

// 347. 前 K 个高频元素
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}

	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 394. 字符串解码
func decodeString(s string) string {
	var stk []string
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else { // ']'
			// 跳过']'
			ptr++
			var sub []string
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			// sub从stk中取出，顺序反了，进行调整
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			// 跳过'['
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			// 跳过数字
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

// 399. 除法求值 (有问题)
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var all = make(map[string]struct{})
	var vs = make(map[string]float64)
	for i, equation := range equations {
		var f0, f1 = true, true
		for i2, s := range equation {
			if _, ok := all[s]; !ok {
				if i2 == 0 {
					f0 = false
				}

				if i2 == 1 {
					f1 = false
				}
				all[s] = struct{}{}
			}
		}

		if f0 && f1 {
			continue
		}

		if !f0 && !f1 {
			vs[equation[0]] = 1.0
			vs[equation[1]] = 1.0 / values[i]
		}

		if f0 {
			vs[equation[1]] = vs[equation[0]] / values[i]
		}

		if f1 {
			vs[equation[0]] = vs[equation[1]] * values[i]
		}
	}

	var ans []float64
	for _, query := range queries {
		var f0, f1 = true, true
		for i2, s := range query {
			if _, ok := all[s]; !ok {
				if i2 == 0 {
					f0 = false
				}

				if i2 == 1 {
					f1 = false
				}
			}
		}

		if !f0 || !f1 {
			ans = append(ans, -1.0)
			continue
		}

		ans = append(ans, vs[query[0]]/vs[query[1]])
	}

	return ans
}

// 399. 除法求值
func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	u := &unionFind{}
	u.init(2*n + 1)
	m := make(map[string]int)
	id := 1
	for i := 0; i < n; i++ {
		s1 := equations[i][0]
		s2 := equations[i][1]
		if _, ok := m[s1]; !ok {
			m[s1] = id
			id++
		}
		if _, ok := m[s2]; !ok {
			m[s2] = id
			id++
		}
		u.union(m[s1], m[s2], values[i])
	}

	n1 := len(queries)
	ans := make([]float64, n1)
	for i := 0; i < n1; i++ {
		s1 := queries[i][0]
		s2 := queries[i][1]
		id1 := m[s1]
		id2 := m[s2]

		if id1 == 0 || id2 == 0 {
			ans[i] = -1.0
		} else {
			ans[i] = u.isConnected(id1, id2)
		}
	}
	return ans
}

type unionFind struct {
	parent []int
	weight []float64
}

func (u *unionFind) init(n int) {
	u.parent = make([]int, n)
	u.weight = make([]float64, n)

	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.weight[i] = 1.0
	}
}

func (u *unionFind) union(x, y int, value float64) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}

	u.parent[rootX] = rootY
	u.weight[rootX] = u.weight[y] * value / u.weight[x]
}

func (u *unionFind) find(x int) int {
	if x != u.parent[x] {
		origin := u.parent[x]
		u.parent[x] = u.find(u.parent[x])
		u.weight[x] *= u.weight[origin]
	}

	return u.parent[x]
}

func (u *unionFind) isConnected(x, y int) float64 {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return u.weight[x] / u.weight[y]
	}

	return -1.0
}
