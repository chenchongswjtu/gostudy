package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	//fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	//fmt.Println(search([]int{1}, 3))
	//fmt.Println(search([]int{5, 1, 3}, 5))
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))

	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))

	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

// 二分查找
// [left,right]
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// 查找左边界
// [left,right]
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1 // 向左缩
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if left > len(nums)-1 || nums[left] != target {
		return -1
	}

	return left
}

// 查找右边界
// [left,right]
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1 // 向右缩
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if right <= 0 || nums[right] != target {
		return -1
	}

	return right
}

// 二分查找
// 珂珂吃香蕉
// 875
// 速度为x时，吃完所有的香蕉需要的时间
// 单调递减函数
func f(piles []int, x int) int {
	var hours int
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[i]%x > 0 {
			hours++
		}
	}
	return hours
}

// 二分查找时间为H的最小速度
// 二分查找左侧
func minEatingSpeed(piles []int, H int) int {
	left := 1         //最小速度
	right := piles[0] // 最大速度
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	for left <= right {
		mid := left + (right-left)/2
		if f(piles, mid) == H {
			right = mid - 1
		} else if f(piles, mid) > H {
			left = mid + 1
		} else if f(piles, mid) < H {
			right = mid - 1
		}
	}

	return left
}

// x 运载能力
// 多少天能运完
// 递减函数
func f1(weights []int, x int) int {
	var days = 0
	for i := 0; i < len(weights); {
		capacity := x
		for i < len(weights) {
			if capacity < weights[i] {
				break
			}
			capacity -= weights[i]
			i++
		}
		days++
	}
	return days
}

// 870 田忌赛马
// 将⻬王和⽥忌的⻢按照战⽃⼒排序，然后按照排名⼀⼀对⽐。如果⽥忌的⻢能赢，那就⽐赛，如果赢不了，
// 那就换个垫底的来送⼈头，保存实⼒。
func advantageCount(nums1, nums2 []int) []int {
	type pair struct {
		i int
		v int
	}

	var pairs []pair
	for i, n := range nums2 {
		pairs = append(pairs, pair{i, n})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].v < pairs[j].v
	})

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums2[j]
	})

	left := 0
	right := len(nums1) - 1
	res := make([]int, len(nums1))

	for i := len(pairs) - 1; i >= 0; i-- {
		index := pairs[i].i
		v := pairs[i].v

		if v < nums1[right] {
			res[index] = nums1[right]
			right--
		} else {
			res[index] = nums1[left]
			left++
		}
	}

	return res
}
