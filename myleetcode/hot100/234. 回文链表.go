package main

//234. 回文链表
//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var nums []int
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next

	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}

	return true
}
