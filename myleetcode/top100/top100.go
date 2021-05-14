package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchRange1([]int{5, 7, 7, 8, 8, 10}, 8))
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
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
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

		if nums[m] >= target {
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
