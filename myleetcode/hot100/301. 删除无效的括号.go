package main

//给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
//
//返回所有可能的结果。答案可以按 任意顺序 返回。

//示例 1：
//
//输入：s = "()())()"
//输出：["(())()","()()()"]
//示例 2：
//
//输入：s = "(a)())()"
//输出：["(a())()","(a)()()"]
//示例 3：
//
//输入：s = ")("
//输出：[""]

func removeInvalidParentheses(s string) []string {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}

	var ans []string
	helper(&ans, s, 0, lremove, rremove)
	return ans
}

func helper(ans *[]string, s string, start int, lremove int, rremove int) {
	if lremove == 0 && rremove == 0 {
		if isValid1(s) {
			*ans = append(*ans, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] { // 去重
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if lremove+rremove > len(s)-i { // 剪枝
			return
		}
		// 尝试去掉一个左括号
		if lremove > 0 && s[i] == '(' {
			helper(ans, s[:i]+s[i+1:], i, lremove-1, rremove) // 删除了i，start还是从i开始
		}
		// 尝试去掉一个右括号
		if rremove > 0 && s[i] == ')' {
			helper(ans, s[:i]+s[i+1:], i, lremove, rremove-1) // 删除了i，start还是从i开始
		}
	}
}

func isValid1(str string) bool {
	count := 0
	for _, ch := range str {
		if ch == '(' {
			count++
		} else if ch == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}
