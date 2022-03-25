package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numStrs := strings.Split(scanner.Text(), " ")
	if len(numStrs) < 1 || len(numStrs) > 1000 {
		return
	}
	var nums = make([]int, len(numStrs))
	for i, str := range numStrs {
		v, err := strconv.Atoi(str)
		if err != nil {
			return
		}
		nums[i] = v
	}
	//fmt.Println(nums)

	var stack []int
	for i := 0; i < len(nums); i++ {
		sum := 0
		j := len(stack) - 1
		for ; j >= 0; j-- {
			if sum+stack[j] == nums[i] {
				sum = sum + stack[j]

				stack = stack[:j]
				stack = append(stack, nums[i]*2)

			REPEAT:
				t := stack[len(stack)-1]
				sum1 := 0
				for k := len(stack) - 2; k >= 0; k-- {
					if sum1+stack[k] == t {
						stack = stack[:k]
						stack = append(stack, t+t)
						goto REPEAT
					} else if sum1+stack[k] > t {
						break
					} else {
						sum1 = sum1 + stack[k]
					}
				}
				break
			} else if sum+stack[j] > nums[i] {
				sum = sum + stack[j]
				stack = append(stack, nums[i])
				break
			} else { // sum+stack[j] < nums[i]
				sum = sum + stack[j]
			}
		}
		if sum < nums[i] {
			stack = append(stack, nums[i])
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%d", stack[i])
		} else {
			fmt.Printf("%d ", stack[i])
		}
	}
}
