package main

//5. 最长回文子串
//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
//示例 1：
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//示例 2：
//
//输入：s = "cbbd"
//输出："bb"
//示例 3：
//
//输入：s = "a"
//输出："a"
//示例 4：
//
//输入：s = "ac"
//输出："a"
//
//
//提示：
//
//1 <= s.length <= 1000
//s 仅由数字和英文字母（大写和/或小写）组成

// 中心扩散(无注记符)
func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	maxLen := 1
	left, right := 0, 0

	for i := 0; i < length; i++ {
		for l, r := i, i; l >= 0 && r < length; {
			if s[l] != s[r] {
				break
			}
			if maxLen < (r - l + 1) {
				maxLen = r - l + 1
				left = l
				right = r
			}
			l--
			r++
		}
	}

	for i := 0; i < length-1; i++ {
		for l, r := i, i+1; l >= 0 && r < length; {
			if s[l] != s[r] {
				break
			}
			if maxLen < (r - l + 1) {
				maxLen = r - l + 1
				left = l
				right = r
			}
			l--
			r++
		}
	}

	return s[left : right+1]
}
