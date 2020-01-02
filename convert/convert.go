package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string to int
	i, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("string to int err:", err)
	}
	fmt.Println(i)

	// string to int64
	i64, err := strconv.ParseInt("678", 10, 64)
	if err != nil {
		fmt.Println("string to int err:", err)
	}
	fmt.Println(i64)

	// int to string
	s := strconv.Itoa(int(67))
	fmt.Println(s)

	// int64 to string
	s = strconv.FormatInt(int64(67), 10)
	fmt.Println(s)
}
