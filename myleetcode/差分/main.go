package main

import (
	"fmt"
)

func main() {
	nums := []int{6, 3, 2, 2}
	diff := newDiff(nums)
	diff.inc(1, 2, 3)
	fmt.Println(diff.result())

}

type diff struct {
	data *[]int
}

func newDiff(nums []int) *diff {
	if len(nums) == 0 {
		return nil
	}

	d := make([]int, len(nums))
	d[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		d[i] = nums[i] - nums[i-1]
	}

	return &diff{data: &d}

}

func (d *diff) inc(i, j, v int) {
	if i > j {
		return
	}
	(*d.data)[i] += v
	if j+1 < len(*d.data) {
		(*d.data)[j+1] -= v
	}
}

func (d *diff) result() []int {
	res := make([]int, len(*d.data))
	res[0] = (*d.data)[0]
	for i := 1; i < len(*d.data); i++ {
		res[i] = res[i-1] + (*d.data)[i]
	}
	return res
}
