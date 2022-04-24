package main

import (
	"container/heap"
	"math/rand"
	"time"
)

// 使用小根堆
// 还可以使用二分排序查找第n-k个数

//215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	var h = &minHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 快速排序
func findKthLargest1(nums []int, k int) int {
	shuffle(nums)
	k1 := len(nums) - k
	l := 0
	r := len(nums) - 1
	for l <= r {
		p := partition(nums, l, r)
		if k1 < p {
			r = p - 1
		} else if k1 > p {
			l = p + 1
		} else {
			return nums[p]
		}
	}
	return -1
}

func partition(nums []int, l int, r int) int {
	v := nums[l]
	i := l + 1
	j := r
	for i <= j {
		for i < r && nums[i] <= v {
			i++
		}
		for j > l && nums[j] >= v {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

// 洗牌算法
// 打乱顺序
func init() {
	rand.Seed(time.Now().UnixNano())
}
func shuffle(nums []int) {
	for i := 0; i < len(nums); i++ {
		t := i + rand.Intn(len(nums)-i)
		nums[i], nums[t] = nums[t], nums[i]
	}
}
