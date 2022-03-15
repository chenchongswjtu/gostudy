package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if len(text) == 0 {
			break
		}
		nums := strings.Split(text, " ")
		var sum int
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
				continue
			}
			sum += n
		}
		fmt.Println(sum)
	}
}
