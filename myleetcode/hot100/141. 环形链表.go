package main

// 哈希表
func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := visited[head]; ok {
			return true
		}
		visited[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
