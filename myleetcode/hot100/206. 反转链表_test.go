package main

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	nl := reverseList(ConvertSliceToListNode([]int{1, 2, 3, 4, 5}))
	p := nl
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}

	nl1 := reverseList1(ConvertSliceToListNode([]int{1, 2, 3, 4, 5}))
	p1 := nl1
	for p1 != nil {
		fmt.Println(p1.Val)
		p1 = p1.Next
	}
}
