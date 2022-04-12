package main

//142. 环形链表 II
//给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//不允许修改 链表。
// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 相等说明存在环，之后fast从head开始，slow继续，再相遇的地方就是环的起点
			fast = head
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

// hashmap，将访问过的存在map中
func detectCycle1(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}
