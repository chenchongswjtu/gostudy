package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

// 3. 无重复字符的最长子串(滑动窗口)
func lengthOfLongestSubstring(s string) int {
	maxLen, curLen := 0, 0
	set := make(map[uint8]struct{})
	var left = 0
	for i := 0; i < len(s); i++ {
		_, ok := set[s[i]]
		for ok {
			delete(set, s[left])
			left++
			curLen--
			_, ok = set[s[i]]
		}

		set[s[i]] = struct{}{}
		curLen++
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 {
		if n2%2 == 0 {
			return float64(nums2[(n2-1)/2]+nums2[(n2-1)/2+1]) / 2
		} else {
			return float64(nums2[n2/2])
		}
	}

	if n2 == 0 {
		if n1%2 == 0 {
			return float64(nums1[(n1-1)/2]+nums1[(n1-1)/2+1]) / 2
		} else {
			return float64(nums1[n1/2])
		}
	}

	if (n1+n2)%2 == 0 {
		target1, target2 := (n1+n2-1)/2+1, (n1+n2-1)/2+2
		t1, t2 := -1, -1
		count := 0
		i, j := 0, 0
		ans := 0
		for i < n1 || j < n2 {
			if i < n1 && j < n2 {
				if nums1[i] <= nums2[j] {
					ans = nums1[i]
					i++
					count++
				} else {
					ans = nums2[j]
					j++
					count++
				}
			} else if i < n1 {
				ans = nums1[i]
				i++
				count++
			} else if j < n2 {
				ans = nums2[j]
				j++
				count++
			}

			if count == target1 {
				t1 = ans
			}

			if count == target2 {
				t2 = ans
			}

			if t1 != -1 && t2 != -1 {
				return float64(t1+t2) / 2
			}
		}
	} else {
		target := (n1+n2)/2 + 1
		count := 0
		i, j := 0, 0
		ans := 0
		for i < n1 || j < n2 {
			if i < n1 && j < n2 {
				if nums1[i] <= nums2[j] {
					ans = nums1[i]
					i++
					count++
				} else {
					ans = nums2[j]
					j++
					count++
				}
			} else if i < n1 {
				ans = nums1[i]
				i++
				count++
			} else if j < n2 {
				ans = nums2[j]
				j++
				count++
			}
			if count == target {
				return float64(ans)
			}
		}
	}
	return -1
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if (n1+n2)&1 == 0 {
		l := find(nums1, 0, nums2, 0, (n1+n2)/2)
		r := find(nums1, 0, nums2, 0, (n1+n2)/2+1)
		return float64(l+r) / 2
	}

	return float64(find(nums1, 0, nums2, 0, (n1+n2)/2+1))
}

func find(nums1 []int, i int, nums2 []int, j int, k int) int {
	if len(nums1)-i > len(nums2)-j {
		return find(nums2, j, nums1, i, k)
	}

	if len(nums1) == i {
		return nums2[j+k-1]
	}

	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	si := min(len(nums1), i+k/2)
	sj := j + k - k/2
	if nums1[si-1] > nums2[sj-1] {
		return find(nums1, i, nums2, sj, k-(sj-j))
	} else {
		return find(nums1, si, nums2, j, k-(si-i))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 11. 盛最多水的容器
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		a := min(height[l], height[r]) * (r - l)
		ans = max(ans, a)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
