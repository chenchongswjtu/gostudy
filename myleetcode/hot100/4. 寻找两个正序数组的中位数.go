package main

//4. 寻找两个正序数组的中位数
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
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
//示例 3：
//
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//示例 4：
//
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
//示例 5：
//
//输入：nums1 = [2], nums2 = []
//输出：2.00000
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
//
//进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
// 二分法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if (n1+n2)%2 == 0 { // 偶数
		l := find(nums1, 0, nums2, 0, (n1+n2)/2)
		r := find(nums1, 0, nums2, 0, (n1+n2)/2+1)
		return float64(l+r) * 0.5
	}
	// 奇数
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
	sj := j + k - k/2 // 必须使用k-k/2，不能使用k/2，会有一个数的差距

	if nums1[si-1] < nums2[sj-1] {
		return find(nums1, si, nums2, j, k-(si-i))
	} else {
		return find(nums1, i, nums2, sj, k-(sj-j))
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
