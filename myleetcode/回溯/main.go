package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
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
