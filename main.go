package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//fmt.Println(isValid1("()[]{}"))
	//fmt.Println(threeSum1([]int{-1, 0, 1, 2, -1, -4}))

	fmt.Println(sortColors([]int{2, 0, 2, 1, 1, 0}))
}

func threeSum1(nums []int) [][]int {
	n := len(nums)
	ret := make([][]int, 0)
	if n < 3 {
		return ret
	}

	set := make(map[string]struct{})
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}

		l, r := i+1, n-1
		for {
			if l >= r {
				break
			}

			if nums[i]+nums[l]+nums[r] == 0 {
				k := strconv.Itoa(nums[i]) + "#" + strconv.Itoa(nums[l]) + "#" + strconv.Itoa(nums[r])
				if _, ok := set[k]; !ok {
					ret = append(ret, []int{nums[i], nums[l], nums[r]})
					set[k] = struct{}{}
				}
				l++ // 相等要进行指针的移动
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}

func threeSum(nums []int) [][]int {
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

func sortColors(nums []int) []int {
	n := len(nums)
	p0, p1 := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p1++
			p0++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}

	return nums
}

// 快速排序
func sortArray(nums []int) []int {
	sort1(nums, 0, len(nums)-1)
	return nums
}

func sort1(nums []int, l, r int) {
	if l >= r {
		return
	}

	m := partition(nums, l, r)
	sort1(nums, l, m-1)
	sort1(nums, m+1, r)
}

func partition(nums []int, l int, r int) int {
	v := nums[l]
	i := l + 1
	j := r
	for i <= j { // i == j 有两个数，还是需要排序的
		for i < r && nums[i] <= v { // i < r ,保证i++之后还是在范围内
			i++
		}

		for j > l && nums[j] >= v { // i < l ,保证j++之后还是在范围内
			j--
		}

		if i >= j { // i == j 不用继续执行，不用交换
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l] // j 指向的数是小于v，交换，最后返回j的值
	return j
}
