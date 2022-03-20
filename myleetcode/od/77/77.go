package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(a77([]int{23, 26, 36, 27}, 78))
	fmt.Println(a77([]int{23, 30, 40}, 26))
}

func a77(nums []int, total int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	k := i + 1
	var ret = -1
	for i < k && k < j {
		if nums[i]+nums[j]+nums[k] == total {
			return total
		} else if nums[i]+nums[j]+nums[k] > total {
			j--
		} else { // nums[i]+nums[j]+nums[k] < total
			for k = i + 1; k < j; k++ {
				if nums[i]+nums[j]+nums[k] > total {
					break
				}
				if nums[i]+nums[j]+nums[k] == total {
					return total
				} else { // nums[i]+nums[j]+nums[k] < total
					if nums[i]+nums[j]+nums[k] > ret {
						ret = nums[i] + nums[j] + nums[k]
					}
				}

			}
			i++
			k = i + 1
		}
	}

	return ret
}
