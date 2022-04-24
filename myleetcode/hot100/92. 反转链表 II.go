package main

// 92. 反转链表 II
//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var dummy = &ListNode{}
	dummy.Next = head //虚拟头节点

	var p = dummy
	var index = 0
	var pre *ListNode
	var start *ListNode
	var end *ListNode
	var next *ListNode
	for p != nil {
		if index == left-1 {
			pre = p
		} else if index == left {
			start = p
		} else if index == right {
			end = p
		} else if index == right+1 {
			next = p
		}
		p = p.Next
		index++
	}

	pre.Next = nil
	end.Next = nil // 将链表切成3段

	newHead := reverse1(start) //翻转中间这一段链表
	pre.Next = newHead
	start.Next = next
	return dummy.Next
}

func reverse1(start *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	var cur = start
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
