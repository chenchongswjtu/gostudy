package main

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//示例 2：
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//提示：
//
//nums1.length == m
//nums2.length == n
//0 <= m <= 1000
//0 <= n <= 1000
//1 <= m + n <= 2000
//-106 <= nums1[i], nums2[i] <= 106
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if (n1+n2)%2 == 0 {
		l := find(nums1, 0, nums2, 0, (n1+n2)/2)
		r := find(nums1, 0, nums2, 0, (n1+n2)/2+1)
		return float64(l+r) / 2.0
	}

	return float64(find(nums1, 0, nums2, 0, (n1+n2)/2+1))
}

// 从nums1数组的第i个开始，从nums2数组的第j个开始，查找第k小的的数（k从1开始）
func find(nums1 []int, i int, nums2 []int, j int, k int) int {
	if len(nums1)-i > len(nums2)-j { // 将短的放到前面
		return find(nums2, j, nums1, i, k)
	}

	if len(nums1) == i { // i到短的数组的最后的，只需要在nums2中找
		return nums2[j+k-1]
	}

	if k == 1 { // 在两个数组中选最小的
		return min(nums1[i], nums2[j])
	}
	si := min(len(nums1), i+k/2) // 找i+k/2的索引，因为nums短，可能不够，取最大的值
	sj := j + k - k/2
	if nums1[si-1] < nums2[sj-1] {
		return find(nums1, si, nums2, j, k-(si-i)) // 将i到si-1的舍去
	} else {
		return find(nums1, i, nums2, sj, k-(sj-j)) // 将j到sj-1舍去
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
