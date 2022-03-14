package main

//17. 电话号码的字母组合
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
//示例 1：
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//示例 2：
//
//输入：digits = ""
//输出：[]
//示例 3：
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//提示：
//
//0 <= digits.length <= 4
//digits[i] 是范围 ['2', '9'] 的一个数字。

// 生成数字和字符串的对应关系map
var m = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// 模拟
func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		str := m[digits[i]]
		tmp := make([]string, 0)
		if i == 0 {
			for j := 0; j < len(str); j++ {
				tmp = append(tmp, string(str[j]))
			}
		} else {
			for k := 0; k < len(ret); k++ {
				for j := 0; j < len(str); j++ {
					tmp = append(tmp, ret[k]+string(str[j]))
				}
			}
		}
		ret = tmp
	}
	return ret
}

// 回溯
func letterCombinations1(digits string) []string {
	var ret []string
	letterCombinations1Helper(digits, 0, "", &ret)
	return ret
}

func letterCombinations1Helper(digits string, index int, s string, ret *[]string) {
	if index == len(digits) {
		*ret = append(*ret, s)
		return
	}

	str := m[digits[index]]
	for i := 0; i < len(str); i++ {
		letterCombinations1Helper(digits, index+1, s+string(str[i]), ret)
	}
}
