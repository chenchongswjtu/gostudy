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

	// uin664 to string
	fmt.Println(strconv.FormatUint(34, 10))

	// float64 to string
	fmt.Println(strconv.FormatFloat(123.66, 'f', 6, 64))

	// float32 to string
	fmt.Println(strconv.FormatFloat(123.66, 'f', -1, 32))

	// string to float32
	f32, err := strconv.ParseFloat("+44.3", 32)
	if err != nil {
		fmt.Println("string to float32 err:", err)
	}
	fmt.Println(f32)
	// 保留两位小数，并且会四舍五入
	fmt.Printf("%.2f \n", f32)

	// string to float64
	f64, err := strconv.ParseFloat("-44.3", 64)
	if err != nil {
		fmt.Println("string to float32 err:", err)
	}
	fmt.Println(f64)
}
