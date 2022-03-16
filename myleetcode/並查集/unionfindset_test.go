package main

import (
	"fmt"
	"testing"
)

func TestUnionFindSet(t *testing.T) {
	uf := NewUnionFindSet(4)
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(1, 3)
	for i := 0; i < 4; i++ {
		fmt.Println(uf.Find(i))
	}
}
