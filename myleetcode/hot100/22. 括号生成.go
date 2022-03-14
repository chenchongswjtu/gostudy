package main

//22. 括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例 1：
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//示例 2：
//
//输入：n = 1
//输出：["()"]
//
//
//提示：
//
//1 <= n <= 8

// 根据“剩余左括号总数要小于等于右括号”的规则生成所有可能结果 + 递归实现
func generateParenthesis(n int) []string {
	var all []string
	generateParenthesisHelper(n, 0, 0, "", &all)
	return all
}

func generateParenthesisHelper(n int, l int, r int, one string, all *[]string) {
	if l == n && r == n {
		*all = append(*all, one)
		return
	}

	if l < r {
		return
	}

	if l < n {
		generateParenthesisHelper(n, l+1, r, one+"(", all)
	}

	if r < n {
		generateParenthesisHelper(n, l, r+1, one+")", all)
	}
}
