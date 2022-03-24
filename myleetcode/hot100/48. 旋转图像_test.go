package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	fmt.Println(rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
