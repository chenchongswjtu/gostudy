package main

// 160.相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)

	pa := headA
	for pa != nil {
		visited[pa] = true
		pa = pa.Next
	}

	pb := headB
	for pb != nil {
		if visited[pb] {
			return pb
		}
		pb = pb.Next
	}

	return nil
}
