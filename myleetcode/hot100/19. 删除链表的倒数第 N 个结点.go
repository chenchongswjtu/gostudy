package main

//19. 删除链表的倒数第 N 个结点
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//进阶：你能尝试使用一趟扫描实现吗？
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left := head
	right := head

	for i := 0; i < n; i++ {
		if right == nil {
			return head
		}
		right = right.Next
	}

	for {
		if right == nil {
			return head.Next
		} else if right.Next != nil {
			right = right.Next
			left = left.Next
		} else if right.Next == nil {
			break
		}
	}
	left.Next = left.Next.Next
	return head
}
