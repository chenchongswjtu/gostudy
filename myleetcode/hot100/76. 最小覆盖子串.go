package main

import (
	"math"
)

//76. 最小覆盖子串
//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//示例 2：
//
//输入：s = "a", t = "a"
//输出："a"
//
//
//提示：
//
//1 <= s.length, t.length <= 105
//s 和 t 由英文字母组成
//
//
//进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

////模板
///* 滑动窗口算法框架 */
//void slidingWindow(string s, string t) {
//Map<Character, Integer> need = new HashMap<>();
//Map<Character, Integer> window = new HashMap<>();
//for (char c : t.toCharArray())
//need.put(c,need.getOrDefault(c,0)+1);
//int left = 0, right = 0;
//int valid = 0;
//while (right < s.size()) {
//// c 是将移入窗口的字符
//char c = s.charAt(right);
//// 右移窗口
//right++;
//// 进行窗口内数据的一系列更新
//...
//
///*** debug 输出的位置 ***/
//System.out.println("window: ["+left+","+ right+")");
///********************/
//
//// 判断左侧窗口是否要收缩
//while (window needs shrink) {
//// d 是将移出窗口的字符
//char d = s[left];
//// 左移窗口
//left++;
//// 进行窗口内数据的一系列更新
//...
//}
//}
//}

func minWindow(s string, t string) string {
	//1.维护两个map记录窗口中的符合条件的字符以及need的字符
	window := make(map[byte]int)
	need := make(map[byte]int) //need中存储的是需要的字符以及需要的对应的数量

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0   //双指针[left,right)
	count := 0            //count记录当前窗口中符合need要求的字符的数量,当count == need.size()时即可shrik窗口
	start := 0            //start表示符合最优解的substring的起始位序
	leng := math.MaxInt64 //len用来记录最终窗口的长度，并且以len作比较，淘汰选出最小的substring的len

	for right < len(s) {
		c := s[right]
		right++ // 窗口扩大
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				count++
			}
		}

		for count == len(need) {
			if right-left < leng {
				leng = right - left
				start = left
			}

			d := s[left]
			left++ // 窗口缩小
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					count--
				}
				window[d]--
			}
		}
	}

	if leng == math.MaxInt64 {
		return ""
	}

	return s[start : start+leng]
}
