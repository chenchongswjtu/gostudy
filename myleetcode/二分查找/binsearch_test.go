package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tables := []struct {
		nums     []int
		target   int
		excepted int
	}{
		{
			nums:     []int{1, 2, 2, 2, 5, 6},
			target:   2,
			excepted: 2,
		},
	}

	for _, tt := range tables {
		actual := binarySearch(tt.nums, tt.target)
		assert.Equal(t, tt.excepted, actual, "")
	}
}

func TestLeftBound(t *testing.T) {
	tables := []struct {
		nums     []int
		target   int
		excepted int
	}{
		{
			nums:     []int{1, 2, 2, 2, 5, 6},
			target:   2,
			excepted: 1,
		},
	}

	for _, tt := range tables {
		actual := leftBound(tt.nums, tt.target)
		assert.Equal(t, tt.excepted, actual, "")
	}
}

func TestRightBound(t *testing.T) {
	tables := []struct {
		nums     []int
		target   int
		excepted int
	}{
		{
			nums:     []int{1, 2, 2, 2, 5, 6},
			target:   2,
			excepted: 3,
		},
		{
			nums:     []int{1, 2, 2, 2, 5, 6},
			target:   0,
			excepted: -1,
		},
		{
			nums:     []int{1, 2, 2, 2, 5, 6},
			target:   7,
			excepted: -1,
		},
	}

	for _, tt := range tables {
		actual := rightBound(tt.nums, tt.target)
		assert.Equal(t, tt.excepted, actual, "")
	}
}
