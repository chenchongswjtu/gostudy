package main

import "fmt"

// 给一个栈的压入顺序和弹出顺序，判断弹出顺序是不是对的
func isValidate(pushed, popped []int) bool {
	stack := make([]int, 0)
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isValidate([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(isValidate([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
