package main

import (
	"fmt"
)

func main() {
	fmt.Println(List2Ints(reverse(Ints2List([]int{1, 2, 3, 4, 5, 6}))))
	fmt.Println(List2Ints(reverse1(Ints2List([]int{1, 2, 3, 4, 5, 6}))))
	fmt.Println(List2Ints(reverseN(Ints2List([]int{1, 2, 3, 4, 5, 6}), 3)))
	fmt.Println(List2Ints(reverseBetween(Ints2List([]int{1, 2, 3, 4, 5, 6}), 2, 4)))
	fmt.Println(List2Ints(reverseKGroup(Ints2List([]int{1, 2, 3, 4, 5, 6, 7}), 4)))
}

// 非递归反转链表
func reverse(head *listNode) *listNode {
	var pre *listNode
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
func reverse1(head *listNode) *listNode {
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
var successor *listNode

// 递归反转链表前n个节点
func reverseN(head *listNode, n int) *listNode {
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
func reverseBetween(head *listNode, m, n int) *listNode {
	if m == 1 {
		return reverseN(head, n)
	}

	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

// k个为一组反转链表
func reverseKGroup(head *listNode, k int) *listNode {
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
func reverseAB(a, b *listNode) *listNode {
	var pre *listNode
	var cur = a
	for cur != b {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
