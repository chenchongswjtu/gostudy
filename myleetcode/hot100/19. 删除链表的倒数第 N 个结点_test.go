package main

import (
	"testing"
)

func TestTRemoveNthFromEnd(t *testing.T) {
	removeNthFromEnd(ConvertSliceToListNode([]int{1, 2, 3, 4}), 1)
}
