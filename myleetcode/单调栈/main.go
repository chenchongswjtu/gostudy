package main

import "fmt"

func main() {
	//fmt.Println(nextGreaterElement([]int{2, 1, 2, 4, 3}))
	fmt.Println(removeDuplicateLetters("bcabc"))
}

// 单调栈
func nextGreaterElement(nums []int) []int {
	var res = make([]int, len(nums))
	var stack []int
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res
}

// 去除重复字母
// 需保证返回的结果的字典序最小（要求不能打乱其它字符的相对位置）
func removeDuplicateLetters(s string) string {
	var count [256]int
	// 记录每一个字符的个数
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	var stack = make([]int, 0)
	var inStack [256]bool
	for i := 0; i < len(s); i++ {
		// 每次遍历计数减1
		count[s[i]]--
		// 栈中包含则不做操作
		if inStack[s[i]] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > int(s[i]) {
			// 若之后不存在栈顶元素，则停止pop
			if count[stack[len(stack)-1]] == 0 {
				break
			}
			// 若之后还有，则pop
			inStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}

		// 栈顶不包含，入栈
		stack = append(stack, int(s[i]))
		inStack[s[i]] = true
	}
	var res string
	// 将栈中的字符转换为字符串
	for i := 0; i < len(stack); i++ {
		c := fmt.Sprintf("%c", stack[i])
		res += c
	}

	return res
}
