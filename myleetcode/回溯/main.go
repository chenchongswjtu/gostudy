package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(readBinaryWatch(1))
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

// 22. 括号生成
// 暴力解法
func generateParenthesis(n int) []string {
	var ans []string
	genAll("", n, &ans)
	return ans
}

func genAll(s string, n int, ans *[]string) {
	if len(s) == 2*n {
		if isValid(s) {
			*ans = append(*ans, s)
		}
		return
	}
	genAll(s+"(", n, ans)
	genAll(s+")", n, ans)

}

func isValid(s string) bool {
	balance := 0
	for _, c := range s {
		if c == '(' {
			balance++
		}

		if c == ')' {
			balance--
		}

		if balance < 0 {
			return false
		}
	}

	return balance == 0
}

// 22. 括号生成
// 回溯解法
func generateParenthesis1(n int) []string {
	var ans []string
	backtrack("", 0, 0, n, &ans)
	return ans
}

func backtrack(cur string, l int, r int, n int, ans *[]string) {
	if len(cur) == 2*n {
		*ans = append(*ans, cur)
		return
	}

	if l < n {
		t := cur
		cur += "("
		backtrack(cur, l+1, r, n, ans)
		cur = t
	}

	if r < l {
		t := cur
		cur += ")"
		backtrack(cur, l, r+1, n, ans)
		cur = t
	}
}

// 39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	// candidates 中没有重复的数，不用排序和去重
	var ans [][]int
	var res []int
	combinationSumHelper(candidates, target, res, 0, 0, &ans)
	return ans
}

func combinationSumHelper(candidates []int, target int, res []int, sum int, start int, ans *[][]int) {
	if sum == target {
		t := make([]int, len(res))
		copy(t, res)
		*ans = append(*ans, t)
		return
	}

	if sum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		t := make([]int, len(res))
		copy(t, res)
		res = append(res, candidates[i])
		combinationSumHelper(candidates, target, res, sum+candidates[i], i, ans) // 一个数可以重复，还是i
		res = t
	}
}

// 40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	// candidates 可能含有重复的数，先拍序
	sort.Ints(candidates)
	var ans [][]int
	var res []int
	combinationSum2Helper(candidates, target, res, 0, 0, &ans)
	// ans = [[1 2 5] [1 7] [1 6 1] [2 6] [2 1 5] [7 1]]
	return duplicate(ans)
}

func combinationSum2Helper(candidates []int, target int, res []int, sum int, start int, ans *[][]int) {
	if sum == target {
		t := make([]int, len(res))
		copy(t, res)
		*ans = append(*ans, t)
		return
	}

	if sum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		t := make([]int, len(res))
		copy(t, res)
		res = append(res, candidates[i])
		combinationSum2Helper(candidates, target, res, sum+candidates[i], i+1, ans) // 不可用重复i+1
		res = t
	}
}

func duplicate(ans [][]int) [][]int {
	var m = make(map[string][]int)
	for _, a := range ans {
		k := toString(a)
		m[k] = a
	}

	var res = make([][]int, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func toString(nums []int) string {
	var res string
	for _, n := range nums {
		res += strconv.Itoa(n)
	}
	return res
}

// 46. 全排列
func permute(nums []int) [][]int {
	all := make([][]int, 0)
	one := make([]int, 0)
	visited := make([]bool, len(nums))
	permuteHelper(nums, visited, one, &all)
	return all
}

func permuteHelper(nums []int, visited []bool, one []int, all *[][]int) {
	if len(one) == len(nums) {
		t := make([]int, len(one))
		copy(t, one)
		*all = append(*all, one)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		t := make([]int, len(one))
		copy(t, one)
		one = append(one, nums[i])
		visited[i] = true
		permuteHelper(nums, visited, one, all)
		one = t
		visited[i] = false
	}
}

// 47. 全排列 II
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	all := make([][]int, 0)
	one := make([]int, 0)
	visited := make([]bool, len(nums))
	permuteHelper(nums, visited, one, &all)
	return duplicate(all)
}

// 47. 全排列 II
func permuteUnique2(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	var perm []int
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || (i > 0 && v == nums[i-1] && !vis[i-1]) {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

// 77. 组合
// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	var all [][]int
	var one []int
	combineHelper(n, k, one, &all)
	return all
}

func combineHelper(n int, k int, one []int, all *[][]int) {
	if len(one) == k {
		t := make([]int, len(one))
		copy(t, one)
		*all = append(*all, one)
		return
	}

	start := 0
	if len(one) == 0 {
		start = 1
	} else {
		start = one[len(one)-1] + 1
	}

	for i := start; i <= n; i++ {
		t := make([]int, len(one))
		copy(t, one)
		one = append(one, i)
		combineHelper(n, k, one, all)
		one = t
	}
}

// 78. 子集
func subsets(nums []int) [][]int {
	var all [][]int
	var one []int
	subsetsHelper(nums, one, &all, 0)
	return all
}

func subsetsHelper(nums []int, one []int, all *[][]int, start int) {
	if start == len(nums) {
		t := make([]int, len(one))
		copy(t, one)
		*all = append(*all, one)
		return
	}

	subsetsHelper(nums, one, all, start+1)
	subsetsHelper(nums, append(one, nums[start]), all, start+1)
}

// 79. 单词搜索
// fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
//func exist(board [][]byte, word string) bool {
//	m := len(board)
//	if m == 0 {
//		return false
//	}
//
//	n := len(board[0])
//	if n == 0 {
//		return false
//	}
//
//	visited := make([][]bool, m)
//	for i := 0; i < m; i++ {
//		visited[i] = make([]bool, n)
//	}
//
//	return existHelper(board, word, visited, 0, 0, m, n, 0)
//}
//
//func existHelper(board [][]byte, word string, visited [][]bool, x, y, m, n, pos int) bool {
//	if pos == len(word) {
//		return true
//	}
//
//	for i := x; i < m; i++ {
//		for j := y; j < n; j++ {
//			if !visited[i][j] && board[i][j] == word[pos] {
//				ans := false
//				visited[i][j] = true
//
//				if i == 0 && j == 0 {
//					ans = ans || existHelper(board, word, visited, i+1, j, m, n, pos+1) ||
//						existHelper(board, word, visited, i, j+1, m, n, pos+1)
//				} else if i == 0 && j == n-1 {
//					ans = ans || existHelper(board, word, visited, i+1, j, m, n, pos+1) ||
//						existHelper(board, word, visited, i, j-1, m, n, pos+1)
//				} else if i == m-1 && j == 0 {
//					ans = ans || existHelper(board, word, visited, i, j+1, m, n, pos+1) ||
//						existHelper(board, word, visited, i-1, j, m, n, pos+1)
//				} else if i == m-1 && j == n-1 {
//					ans = ans || existHelper(board, word, visited, i, j-1, m, n, pos+1) ||
//						existHelper(board, word, visited, i-1, j, m, n, pos+1)
//				} else if i == 0 {
//					ans = ans || existHelper(board, word, visited, i+1, j, m, n, pos+1) ||
//						existHelper(board, word, visited, i, j-1, m, n, pos+1) ||
//						existHelper(board, word, visited, i, j+1, m, n, pos+1)
//				} else if i == m-1 {
//					ans = ans || existHelper(board, word, visited, i-1, j, m, n, pos+1) ||
//						existHelper(board, word, visited, i, j-1, m, n, pos+1) ||
//						existHelper(board, word, visited, i, j+1, m, n, pos+1)
//				} else if j == 0 {
//					ans = ans || existHelper(board, word, visited, i, j+1, m, n, pos+1) ||
//						existHelper(board, word, visited, i-1, j, m, n, pos+1) ||
//						existHelper(board, word, visited, i+1, j, m, n, pos+1)
//				} else if j == n-1 {
//					ans = ans || existHelper(board, word, visited, i, j-1, m, n, pos+1) ||
//						existHelper(board, word, visited, i-1, j, m, n, pos+1) ||
//						existHelper(board, word, visited, i+1, j, m, n, pos+1)
//				} else {
//					ans = ans || existHelper(board, word, visited, i, j-1, m, n, pos+1) ||
//						existHelper(board, word, visited, i, j+1, m, n, pos+1) ||
//						existHelper(board, word, visited, i-1, j, m, n, pos+1) ||
//						existHelper(board, word, visited, i+1, j, m, n, pos+1)
//				}
//
//				if ans {
//					return true
//				}
//
//				visited[i][j] = false
//			}
//		}
//	}
//
//	return false
//}

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
// 79. 单词搜索
func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])

	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() {
			vis[i][j] = false
		}() // 回溯时还原已访问的单元格

		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 89. 格雷编码
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	ans := []int{0}
	head := 1
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]+head)
		}
		head = head << 1
	}

	return ans
}

// 90. 子集 II(先排列所有的子集再去重)
func subsetsWithDup2(nums []int) [][]int {
	sort.Ints(nums)
	var all [][]int
	var one []int
	subsetsHelper(nums, one, &all, 0)
	return duplicate(all)
}

// 90. 子集 II
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var t []int
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1]
	}
	dfs(false, 0)
	return
}

// 93. 复原 IP 地址
func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}
	var all []string
	var one string
	restoreIpAddressesHelper(s, one, &all, 0)
	return all
}

func restoreIpAddressesHelper(s string, one string, all *[]string, cur int) {
	if cur == 4 {
		if len(s) == 0 {
			*all = append(*all, one)
		}
		return
	}

	if len(s) < 4-cur {
		return
	}

	var n1, n2, n3 string
	if len(s) >= 1 {
		n1 = s[:1]
	}

	if len(s) >= 2 {
		n2 = s[:2]
		if n2[0] == '0' {
			n2 = ""
		}
	}

	if len(s) >= 3 {
		n3 = s[:3]
		if n3[0] == '0' {
			n3 = ""
		}
	}

	for i, n := range []string{n1, n2, n3} {
		if len(n) == 0 {
			continue
		}

		i1, _ := strconv.Atoi(n)
		if 0 <= i1 && i1 < 256 {
			t := one
			if len(one) == 0 {
				one = n
			} else {
				one = one + "." + n
			}
			restoreIpAddressesHelper(s[i+1:], one, all, cur+1)
			one = t
		}
	}
}

// 131. 分割回文串
func partition(s string) [][]string {
	var ans [][]string
	var one []string
	partitionHelper(s, one, &ans)
	return ans
}

func partitionHelper(s string, one []string, ans *[][]string) {
	if len(s) == 0 {
		//t := make([]string, len(one))
		//copy(t, one)
		*ans = append(*ans, append([]string(nil), one...))
		return
	}

	var isValid func(n string) bool
	isValid = func(n string) bool {
		i, j := 0, len(n)-1
		for i < j {
			if n[i] != n[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	for i := 0; i < len(s); i++ {
		n := s[:i+1]
		if isValid(n) {
			// copy 一个对象
			//t := make([]string, len(one))
			//copy(t, one)
			one = append(one, n)
			partitionHelper(s[i+1:], one, ans)
			one = one[:len(one)-1]
		}
	}
}

// 回溯加动态规划
func partition2(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var splits []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}

	dfs(0)
	return
}

// 216. 组合总和 III
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	combinationSum3Helper(k, n, 0, 0, 1, []int(nil), &ans)
	return ans
}

func combinationSum3Helper(k int, n int, cur int, sum int, start int, one []int, all *[][]int) {
	if cur == k {
		if sum == n {
			*all = append(*all, append([]int(nil), one...))
		}
		return
	}

	for i := start; i <= 9; i++ {
		one = append(one, i)
		combinationSum3Helper(k, n, cur+1, sum+i, i+1, one, all)
		one = one[:len(one)-1]
	}
}

// 306. 累加数
func isAdditiveNumber(num string) bool {
	return isAdditiveNumberHelper(num, len(num), 0, 0, 0, 0)
}

// @param num    原始字符串
// @param size   原始字符串长度
// @param idx    当前处理下标
// @param sum    前面的两个数字之和
// @param pre    前一个数字
// @param k      当前是处理的第几个数字
func isAdditiveNumberHelper(num string, size int, idx int, sum int, pre int, k int) bool {
	if idx == size {
		return k > 2
	}
	for i := idx; i < size; i++ {
		cur := fetchCurValue(num, idx, i)
		if cur < 0 {
			continue
		}
		if k >= 2 && cur != sum {
			continue
		}
		if isAdditiveNumberHelper(num, size, i+1, pre+cur, cur, k+1) {
			return true
		}
	}
	return false
}

// 获取 l ~ r 组成的有效数字
func fetchCurValue(num string, l int, r int) int {
	if l < r && num[l] == '0' {
		return -1
	}
	res := 0
	for l <= r {
		n := num[l] - '0'
		res = res*10 + int(n)
		l++
	}
	return res
}

// 357. 计算各个位数不同的数字个数
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}

	res := 9
	for i := 1; i < n; i++ {
		res = res * (10 - i)
	}

	return res + countNumbersWithUniqueDigits(n-1)
}

// 401. 二进制手表
func readBinaryWatch(turnedOn int) []string {
	if turnedOn > 10 || turnedOn < 0 {
		return nil
	}

	var ans []string
	for i := 0; i <= 4; i++ {
		j := turnedOn - i
		if j > 6 || j < 0 {
			continue
		}
		ans = append(ans, gen(i, j)...)
	}
	return ans
}

var hm = map[int][]string{
	0: {"0"},
	1: {"1", "2", "4", "8"},
	2: {"3", "5", "9", "6", "10"},
	3: {"7", "11"},
}

var mm = map[int][]string{
	0: {"00"},
	1: {"01", "02", "04", "08", "16", "32"},
	2: {"03", "05", "09", "17", "33", "06", "10", "18", "34", "12", "20", "36", "24", "40", "48"},
	3: {"07", "11", "19", "35", "14", "22", "38", "13", "21", "37", "25", "41", "49", "26", "42", "50", "28", "44", "52", "56"},
	4: {"58", "54", "46", "30", "57", "53", "45", "29", "51", "43", "27", "39", "23", "15"},
	5: {"59", "55", "47", "31"},
}

func gen(h int, m int) []string {
	hv := hm[h]
	mv := mm[m]
	var res []string
	for _, v := range hv {
		for _, vv := range mv {
			res = append(res, v+":"+vv)
		}
	}
	return res
}
