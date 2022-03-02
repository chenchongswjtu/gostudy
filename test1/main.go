package main

import (
	"fmt"
	"time"
)

func main() {
	b := newBlock(size)

	go func() {
		for i := 1; i < 100; i++ {
			b.put(i)
		}
	}()

	for i := 1; i < 10; i++ {
		go b.selector()
	}

	time.Sleep(2 * time.Second)

}

const size = 10

type block struct {
	next int
	nums []int
}

func newBlock(size int) *block {
	return &block{
		next: 1,
		nums: make([]int, size),
	}
}

func (c *block) Get(num int) int {
	index := num % size
	if c.nums[index] == num {
		c.nums[index] = 0
		return num
	}

	return 0
}

func (c *block) selector() {
	n := c.Get(c.next)
	if n != 0 {
		c.next++
		c.selector()
	} else {
		fmt.Printf("Get block %d failed \n", c.next)
	}
}

func (c *block) put(n int) {
	index := n % size
	c.nums[index] = n
}
