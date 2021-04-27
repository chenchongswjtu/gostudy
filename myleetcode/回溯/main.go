package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis1(3))
}

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	letter := make(map[string][]string)
	letter["2"] = []string{"a", "b", "c"}
	letter["3"] = []string{"d", "e", "f"}
	letter["4"] = []string{"g", "h", "i"}
	letter["5"] = []string{"j", "k", "l"}
	letter["6"] = []string{"m", "n", "o"}
	letter["7"] = []string{"p", "q", "r", "s"}
	letter["8"] = []string{"t", "u", "v"}
	letter["9"] = []string{"w", "x", "y", "z"}

	if len(digits) == 0 {
		return nil
	}

	var ans []string
	letterCombinationsHelper(letter, digits, 0, "", &ans)
	return ans
}

func letterCombinationsHelper(letter map[string][]string, digits string, n int, s string, ans *[]string) {
	if n == len(digits) {
		*ans = append(*ans, s)
		return
	}

	for _, v := range letter[digits[n:n+1]] {
		t := s

		s += v // 选择
		letterCombinationsHelper(letter, digits, n+1, s, ans)
		s = t // 撤销选择，回溯
	}
}

// 22. 括号生成
// 暴力解法
func generateParenthesis(n int) []string {
	var ans []string
	genAll("", n, &ans)
	return ans
}

func genAll(s string, n int, ans *[]string) {
	if len(s) == 2*n {
		if isValid(s) {
			*ans = append(*ans, s)
		}
		return
	}
	genAll(s+"(", n, ans)
	genAll(s+")", n, ans)

}

func isValid(s string) bool {
	balance := 0
	for _, c := range s {
		if c == '(' {
			balance++
		}

		if c == ')' {
			balance--
		}

		if balance < 0 {
			return false
		}
	}

	return balance == 0
}

// 22. 括号生成
// 回溯解法
func generateParenthesis1(n int) []string {
	var ans []string
	backtrack("", 0, 0, n, &ans)
	return ans
}

func backtrack(cur string, l int, r int, n int, ans *[]string) {
	if len(cur) == 2*n {
		*ans = append(*ans, cur)
		return
	}

	if l < n {
		t := cur
		cur += "("
		backtrack(cur, l+1, r, n, ans)
		cur = t
	}

	if r < l {
		t := cur
		cur += ")"
		backtrack(cur, l, r+1, n, ans)
		cur = t
	}
}
