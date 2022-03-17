package main

import "fmt"

func main() {
	var nums = []int{3, 2, 41, 4, 2}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// 快速排序
func quickSort(nums []int, low int, high int) {
	if low >= high {
		return
	}
	mid := partition(nums, low, high)
	quickSort(nums, low, mid-1)
	quickSort(nums, mid+1, high)
}

func partition(nums []int, low int, high int) int {
	p := nums[low]
	i, j := low+1, high
	for i <= j { // i==j 也要排序，有两个数
		for i < high && nums[i] <= p {
			i++
		}

		for j > low && nums[j] >= p {
			j--
		}

		if i >= j { // 循环结束
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j], nums[low] = nums[low], nums[j] // 将low与j想交换
	return j
}
