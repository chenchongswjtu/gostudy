package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	numStrs := strings.Split(line, ",")
	if len(numStrs) > 25 {
		return
	}

	for _, str := range numStrs {
		if len(str) > 6 || len(str) < 1 {
			return
		}
	}

	sort.Slice(numStrs, func(i, j int) bool {
		if len(numStrs[i]) == len(numStrs[j]) {
			return numStrs[i] > numStrs[j]
		} else {
			return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
		}
	})

	fmt.Println(strings.Join(numStrs, ""))
}
