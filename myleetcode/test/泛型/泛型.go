package main

import "fmt"

// 泛型
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v \n", v)
	}
}

func main() {
	printSlice[int]([]int{66, 77, 88, 99, 100})
	printSlice[string]([]string{"zhangsan", "lisi", "wangwu", "zhaosi"})
}
