package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tags = []string{"atime", "btime", "ctime", "dtime"}

var sumTag = []int{0, 0, 0, 0}

var sumAll = 0

func readLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file [%s] err: %s", filename, err)
		return
	}

	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("######### sumAll = %d, sumTag = %v ", sumAll, sumTag)
				return
			}
			fmt.Printf("read string err: %s", err)
			return
		}

		// fmt.Println(line)

		start := strings.Index(line, "cost ")
		end := strings.Index(line, " ns")

		fmt.Println(line[start+len("cost ") : end])

		n, err := strconv.Atoi(line[start+len("cost ") : end])
		if err != nil {
			fmt.Printf("open file [%s] err: %s", filename, err)
			return
		}

		fmt.Println(n)

		for i, v := range tags {
			if strings.Index(line, v) >= 0 {
				fmt.Println("-------", i)
				sumTag[i] += n
			}
		}

		sumAll += n

		if sumAll != all(sumTag) {
			fmt.Println("xxxxxxxxx")
		}
	}

}

func all(nums []int) int {
	var sum = 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	readLine("D:\\gopath\\src\\github.com\\chenchongswjtu\\gostudy\\test.txt")
}
