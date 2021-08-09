package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	start = "CommitBlock_Peer "
	end   = " times"
)

func main() {
	all, err := ioutil.ReadFile("D:\\gopath\\src\\github.com\\chenchongswjtu\\gostudy\\mytest\\log.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	count := 0
	lines := strings.Split(string(all), "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}
		l := strings.Index(line, start) + len(start)
		r := strings.Index(line, end)
		num, err := strconv.Atoi(line[l:r])
		if err != nil {
			panic(err)
		}

		fmt.Println(num)
		sum += num
		count++
	}

	avg := sum / count
	fmt.Printf("sum = %d, count = %d, avg = %d \n", sum, count, avg)
}
