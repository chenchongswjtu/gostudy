package main

import (
	"sort"
	"strconv"
)

//15. 三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//示例 2：
//
//输入：nums = []
//输出：[]
//示例 3：
//
//输入：nums = [0]
//输出：[]
//
//
//提示：
//
//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105

// 排序+双指针
// 依次遍历每个元素作为第一元素，双指针分别指向下一个元素与最后一个元素
// 去重：Set去重，(a,b,-(a+b))排序后生成string作为key
// 去重优化：依次遍历每个元素作为第一元素时，重复元素直接跳过；left和right指针与前一个元素相同时也跳过
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	var ret [][]int
	var set = make(map[string]bool)
	for i := 0; i < size-2; i++ {
		if nums[i] > 0 { // 经过排序，最小的大于0，则不用在查找
			return ret
		}

		left, right := i+1, size-1
		for {
			if left >= right {
				break
			}
			if nums[i]+nums[left]+nums[right] == 0 {
				str := strconv.Itoa(nums[i]) + "#" + strconv.Itoa(nums[left]) + "#" + strconv.Itoa(nums[right])
				if _, ok := set[str]; !ok { // map去重
					set[str] = true
					ret = append(ret, []int{nums[i], nums[left], nums[right]})
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ret
}

// 三数之和
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	ret, start, end, index, length, sum := make([][]int, 0), 0, 0, 0, len(nums), 0

	for index = 0; index < length-1; index++ {
		start, end = 0, length-1
		if index >= 1 && nums[index] == nums[index-1] { // 相同，说明已经包含过，将start设置为index-1
			start = index - 1
		}

		for start < index && index < end {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}

			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			sum = nums[start] + nums[index] + nums[end]
			if sum == 0 {
				ret = append(ret, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return ret
}
