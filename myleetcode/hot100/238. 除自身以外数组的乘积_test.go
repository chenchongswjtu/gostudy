package main

import (
	"fmt"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf1([]int{1, 2, 3, 4}))
}
