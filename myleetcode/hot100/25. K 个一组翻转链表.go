package main

// 25.k个一组翻转链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	a := head
	b := head
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

func reverseAB(a *ListNode, b *ListNode) *ListNode {
	var pre *ListNode
	var cur = a
	var next *ListNode

	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
