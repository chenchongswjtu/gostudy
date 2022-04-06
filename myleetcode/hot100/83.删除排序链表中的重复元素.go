package main

// 83.删除排序链表中重复的元素
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	slow := head

	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}

	slow.Next = nil
	return head
}
