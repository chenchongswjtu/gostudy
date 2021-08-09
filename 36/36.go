package main

import (
	"fmt"
)

func main() {
	fmt.Println(add36Strings("5", "9"))
}

func add36Strings(num1, num2 string) string {
	carry := 0
	n1, n2 := len(num1)-1, len(num2)-1
	ans := ""
	x, y := 0, 0
	for n1 >= 0 || n2 >= 0 || carry > 0 {
		if n1 >= 0 {
			x = to36Int(num1[n1])
		} else {
			x = 0
		}

		if n2 >= 0 {
			y = to36Int(num2[n2])
		} else {
			y = 0
		}

		t := x + y + carry
		ans = to36Char(t%36) + ans
		carry = t / 36
		n1--
		n2--
	}
	return ans
}

func to36Int(c uint8) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}

	return int(c-'a') + 10
}

func to36Char(n int) string {
	if n <= 9 {
		return string(rune(n + '0'))
	}

	return string(rune(n - 10 + 'a'))
}
