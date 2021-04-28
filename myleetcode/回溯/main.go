package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
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
	var m = make(map[int][]int)
	for _, a := range ans {
		k := toInt(a)
		m[k] = a
	}

	var res = make([][]int, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func toInt(nums []int) int {
	var sum int
	for _, n := range nums {
		sum = sum*10 + n
	}
	return sum
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
