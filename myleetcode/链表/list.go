package main

import (
	"fmt"
)

func main() {
	//fmt.Println(List2Ints(reverse(Ints2List([]int{1, 2, 3, 4, 5, 6}))))
	//fmt.Println(List2Ints(reverse1(Ints2List([]int{1, 2, 3, 4, 5, 6}))))
	//fmt.Println(List2Ints(reverseN(Ints2List([]int{1, 2, 3, 4, 5, 6}), 3)))
	//fmt.Println(List2Ints(reverseBetween(Ints2List([]int{1, 2, 3, 4, 5, 6}), 2, 4)))
	//fmt.Println(List2Ints(reverseKGroup(Ints2List([]int{1, 2, 3, 4, 5, 6, 7}), 4)))

	fmt.Println(List2Ints(removeNthFromEnd(Ints2List([]int{1, 2, 3, 4, 5}), 2)))
}

// 非递归反转链表
func reverse(head *ListNode) *ListNode {
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

// 递归反转链表
func reverse1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1->2->3-nil
	// 1->reverse1(head-next)(nil-2<-3)
	//							     ^
	// 								 last
	last := reverse1(head.Next)
	// head.Next.Next之前为nil
	head.Next.Next = head

	head.Next = nil
	return last
}

// 递归反转链表前n个节点的后驱节点
var successor *ListNode

// 递归反转链表前n个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 当n为1时记录他的后驱节点
		successor = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

// 递归反转链表前m到n个节点
func reverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}

	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

// k个为一组反转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	var a = head
	var b = head

	for i := k; i > 0; i-- {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverseAB(a, b)
	a.Next = reverseKGroup(b, k)

	return newHead
}

// 非递归反转链表a到b节点
func reverseAB(a, b *ListNode) *ListNode {
	var pre *ListNode
	var cur = a
	for cur != b {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	var carry int
	var head *ListNode
	var cur *ListNode
	for l1 != nil && l2 != nil {
		res := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		node := &ListNode{Val: res, Next: nil}
		if head == nil {
			head = node
			cur = node
		} else {
			cur.Next = node
			cur = cur.Next
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		res := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		node := &ListNode{Val: res, Next: nil}
		cur.Next = node
		cur = cur.Next
		l1 = l1.Next
	}

	for l2 != nil {
		res := (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		node := &ListNode{Val: res, Next: nil}
		cur.Next = node
		cur = cur.Next
		l2 = l2.Next
	}

	if carry != 0 {
		node := &ListNode{Val: carry, Next: nil}
		cur.Next = node
	}

	return head
}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	slow := head

	cur := head
	l := 0
	for cur != nil {
		l++
		cur = cur.Next
	}

	m := l - n
	if m < 0 {
		return head
	}

	if m == 0 {
		return head.Next
	}

	for i := 0; i < m-1; i++ {
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}

	// 删除头节点
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

// 剑指 Offer 18. 删除链表的节点
// 删除所有的val节点
func deleteNode(head *ListNode, val int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head
	prev := dummy
	cur := head
	for cur != nil {
		next := cur.Next
		if cur.Val == val {
			prev.Next = next
			cur = next
			continue
		}
		prev = cur
		cur = next
	}

	return dummy.Next
}
