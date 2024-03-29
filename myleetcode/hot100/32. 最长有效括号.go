package main

//32. 最长有效括号
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//示例 1：
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//
//输入：s = ""
//输出：0
//
//
//提示：
//
//0 <= s.length <= 3 * 104
//s[i] 为 '(' 或 ')'

// 使用(类似)栈优化
func longestValidParentheses(s string) int {
	//算法核心：
	//    始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」，
	//    栈里其他元素维护左括号的下标
	//
	//1. 对于遇到的每个 \text{‘(’}‘(’ ，我们将它的下标放入栈中
	//2. 对于遇到的每个 \text{‘)’}‘)’ ，我们先弹出栈顶元素表示匹配了当前右括号：
	//		(1) 如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的「最后一个没有被匹配的右括号的下标」
	//		(2) 如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
	//Note:
	// 		如果一开始栈为空，第一个字符为左括号的时候我们会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，
	// 		为了保持统一，我们在一开始的时候往栈中放入一个值为 -1−1 的元素。
	size := len(s)
	ret := 0
	if size <= 1 {
		return 0
	}
	stack := []int{-1}

	for i := 0; i < size; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // 第一个为）
				stack = append(stack, i)
			} else {
				ret = maxInt(ret, i-stack[len(stack)-1]) // stack最后一个可能为(
			}
		}
	}
	return ret
}

// 这个方法比较好理解
func longestValidParentheses1(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}
