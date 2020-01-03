package main

import "fmt"

func f() *int {
	var c = 23
	return &c
}

func main() {
	fmt.Println(f())
}
