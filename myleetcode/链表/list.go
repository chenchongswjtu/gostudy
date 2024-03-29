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

// 61. 旋转链表
// 先转为环，再断开
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	// 链表长度
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}

	pos := n - k%n
	if pos == n {
		return head
	}

	cur.Next = head // 合成环
	for pos > 0 {
		cur = cur.Next
		pos--
	}

	res := cur.Next
	cur.Next = nil

	return res
}

// 83. 删除排序链表中的重复元素
// 1->2->3->3->4->4->5
// 1->2->3->4->5
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head.Next

	for {
		for p1.Val == p2.Val {
			p2 = p2.Next
			if p2 == nil {
				p1.Next = nil
				return head
			} else {
				p1.Next = p2
			}
		}
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			return head
		}
	}
}

// 83. 删除排序链表中的重复元素
// 1->2->3->3->4->4->5
// 1->2->3->4->5
// 简单
func deleteDuplicates4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

// 82. 删除排序链表中的重复元素 II
// 1->2->3->3->4->4->5
// 1->2->5
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	p0 := dummy     // 相同的前一个
	p1 := head      // 相同的第一个
	p2 := head.Next // 相同的最后一个

	isSame := false
	for {
		for p1.Val == p2.Val {
			isSame = true
			p2 = p2.Next
			if p2 == nil {
				p0.Next = nil
				return dummy.Next
			}
		}
		if isSame {
			p0.Next = p2
			isSame = false

			p1 = p2
			p2 = p2.Next
			if p2 == nil {
				return dummy.Next
			}
		} else {
			p0 = p0.Next
			p1 = p1.Next
			p2 = p2.Next
			if p2 == nil {
				return dummy.Next
			}
		}
	}
}

func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 86. 分隔链表
func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	cur := head
	for cur != nil {
		if cur.Val < x {
			small.Next = cur
			small = small.Next
		} else {
			large.Next = cur
			large = large.Next
		}
		cur = cur.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

// 92. 反转链表 II
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	cur := dummy
	for i := 1; i < left; i++ {
		cur = cur.Next
	}
	beforeA := cur
	a := cur.Next

	cur = dummy
	for i := 0; i < right; i++ {
		cur = cur.Next
	}
	b := cur.Next // right后一个

	nb := reverseAB1(a, b)

	beforeA.Next = nb
	a.Next = b
	return dummy.Next
}

// [a,b)
func reverseAB1(a, b *ListNode) *ListNode {
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

// 143. 重排链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

// 328. 奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	i, j := 0, 1
	heada := nodes[i]
	headb := nodes[j]
	nodea := heada
	nodeb := headb

	i = i + 2
	j = j + 2
	for i < len(nodes) {
		nodea.Next = nodes[i]
		nodea = nodea.Next
		i = i + 2
	}

	for j < len(nodes) {
		nodeb.Next = nodes[j]
		nodeb = nodeb.Next
		j = j + 2
	}

	nodea.Next = headb
	nodeb.Next = nil
	return heada
}
