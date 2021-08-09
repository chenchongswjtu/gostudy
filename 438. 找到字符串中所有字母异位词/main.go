package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

// 滑动窗口
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	res := make([]int, 0)
	left := 0
	right := 0
	need := make(map[int32]int)
	window := make(map[int32]int)
	match := 0

	for _, c := range p {
		need[c]++
	}

	for right < len(s) {
		c := s[right]
		if _, ok := need[int32(c)]; ok {
			window[int32(c)]++
			if window[int32(c)] == need[int32(c)] {
				match++
			}
		}
		right++

		for match == len(need) {
			if right-left == len(p) {
				res = append(res, left)
			}

			c = s[left]
			if _, ok := need[int32(c)]; ok {
				window[int32(c)]--
				if window[int32(c)] < need[int32(c)] {
					match--
				}
			}
			left++
		}
	}
	return res
}
