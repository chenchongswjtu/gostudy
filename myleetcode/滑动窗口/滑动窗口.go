// 滑动窗口[left, right) 左闭右开
// need为目标字符串的字符的个数
// windows为[left, right)中目标字符的个数

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindows("ebbancf", "abc"))
	fmt.Println(findAnagrams("cbaebabacd", "abc"))

	fmt.Println(lengthOfLongestSubstring("abdsd"))
}

// 求字符串s中包含t的最新字符串
func minWindow(s, t string) string {
	var need = make(map[byte]int)
	var window = make(map[byte]int)

	for _, c := range t {
		need[byte(c)]++
	}

	// [left, right)
	var left = 0
	var right = 0
	var valid = 0         // windows中与need中字符个数相等的个数
	var start = 0         // 开始位置
	var l = math.MaxInt64 // 长度
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < l {
				start = left
				l = right - left
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if l == math.MinInt64 {
		return ""
	}

	return s[start : start+l]
}

// 438
func findAnagrams(s, t string) []int {
	var need = make(map[byte]int)
	var window = make(map[byte]int)

	for _, c := range t {
		need[byte(c)]++
	}

	// [left, right)
	var left = 0
	var right = 0
	var match = 0 // windows中与need中字符个数相等的个数
	var res []int
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				match++
			}
		}
		//for right-left >= len(t) {
		//	if match == len(need) {
		//		res = append(res, left)
		//	}
		//
		//	d := s[left]
		//	left++
		//	if _, ok := need[d]; ok {
		//		if window[d] == need[d] {
		//			match--
		//		}
		//		window[d]--
		//	}
		//}

		for match == len(need) { // 匹配才验证长度
			if right-left == len(t) {
				res = append(res, left)
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					match--
				}
				window[d]--
			}
		}
	}

	return res
}

func minWindows(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	var need = make(map[byte]int)
	var window = make(map[byte]int)
	var left = 0
	var right = 0
	var match = 0
	var length = math.MaxInt64 // 长度初始化为最大
	var start = 0

	for _, c := range t {
		need[byte(c)]++
	}

	for right < len(s) {
		r := s[right]
		right++
		if _, ok := need[r]; ok {
			window[r]++
			if window[r] == need[r] {
				match++
			}
		}

		for match == len(need) {
			if right-left < length {
				length = right - left
				start = left
			}

			l := s[left]
			left++
			if _, ok := need[l]; ok {
				if window[l] == need[l] {
					match--
				}
				window[l]--
			}
		}
	}

	if length == math.MaxInt64 {
		return ""
	}

	return s[start : start+length]
}

// 最长无重复子串
func lengthOfLongestSubstring(s string) int {
	var res = 0
	var left = 0
	var right = 0
	var window = make(map[byte]int)

	for right < len(s) {
		rc := s[right]
		right++
		window[rc]++
		for window[rc] > 1 {
			lc := s[left]
			left++
			window[lc]--
		}

		if right-left > res {
			res = right - left
		}
	}
	return res
}
