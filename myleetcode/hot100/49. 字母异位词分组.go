package main

import "sort"

//49. 字母异位词分组
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//["ate","eat","tea"],
//["nat","tan"],
//["bat"]
//]
//说明：
//
//所有输入均为小写字母。
//不考虑答案输出的顺序。

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	m := make(map[string][]string)

	// 将字符串按照字符从小到大排序作为key
	// 字符串作为value
	for _, str := range strs {
		tmp := str
		bs := []byte(tmp)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})

		m[string(bs)] = append(m[string(bs)], str)
	}

	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}
