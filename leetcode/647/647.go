package main

import (
	"fmt"
)

// 回文子串，中心扩展法
func countSubStrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var init = 0
	var count = &init
	for i := 0; i < len(s); i++ {
		centerSpend(s, i, i, count)
		centerSpend(s, i, i+1, count)
	}
	return *count
}

func centerSpend(s string, l int, r int, count *int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		*count++
		l--
		r++
	}
}

func main() {
	fmt.Println(countSubStrings("aaa"))
}
