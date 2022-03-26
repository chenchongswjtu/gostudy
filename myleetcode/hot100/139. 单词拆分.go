package main

//139. 单词拆分
//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
//说明：
//
//拆分时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false

// 动态规划
func wordBreak(s string, wordDict []string) bool {
	/*
		dp[i]: s[:i]能否拆分成wordDict的组合
		dp[j] = dp[i] && (s[i:j] in wordDict)
	*/
	size := len(s)
	if size == 0 {
		return true
	}

	m := make(map[string]bool, 0)
	for _, word := range wordDict {
		m[word] = true
	}

	dp := make([]bool, size+1)
	dp[0] = true
	for i := 0; i < size; i++ {
		if !dp[i] {
			continue
		}
		for j := i + 1; j <= size; j++ {
			if m[s[i:j]] {
				dp[j] = true
			}
		}
	}
	return dp[size]
}
