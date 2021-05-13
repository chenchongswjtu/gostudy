package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// 3. 无重复字符的最长子串(滑动窗口)
func lengthOfLongestSubstring(s string) int {
	maxLen, curLen := 0, 0
	set := make(map[uint8]struct{})
	var left = 0
	for i := 0; i < len(s); i++ {
		_, ok := set[s[i]]
		for ok {
			delete(set, s[left])
			left++
			curLen--
			_, ok = set[s[i]]
		}

		set[s[i]] = struct{}{}
		curLen++
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
