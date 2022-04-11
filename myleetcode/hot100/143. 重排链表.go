package main

//143. 重排链表
//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//L0 → L1 → … → Ln - 1 → Ln
//请将其重新排列后变为：
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var list []*ListNode
	for p := head; p != nil; p = p.Next {
		list = append(list, p)
	}

	i, j := 0, len(list)-1
	for i < j {
		list[i].Next = list[j]
		i++
		if i == j {
			break
		}
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil // 结束
}
