package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(ConvertSliceToListNode([]int{1, 2, 1})))
}
