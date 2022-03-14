package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, sum int
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := strings.Split(input.Text(), ",")
		sum = 0
		for i := 0; i < len(nums); i++ {
			n, _ = strconv.Atoi(nums[i])
			sum += n
		}
		fmt.Println(sum)
	}
}
