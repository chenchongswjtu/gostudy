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

// 双指针
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa := headA
	pb := headB
	for pa != pb {
		if pa == nil { // headA走完了，接着走headB
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil { // headB走完了，接着走headA
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa // 最后都走完还没有相同的就是nil退出，返回nil
}
