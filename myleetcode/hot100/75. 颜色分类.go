package main

//75. 颜色分类
//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//
//
//示例 1：
//
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
//示例 2：
//
//输入：nums = [2,0,1]
//输出：[0,1,2]
//示例 3：
//
//输入：nums = [0]
//输出：[0]
//示例 4：
//
//输入：nums = [1]
//输出：[1]
//
//
//提示：
//
//n == nums.length
//1 <= n <= 300
//nums[i] 为 0、1 或 2
//
//
//进阶：
//
//你可以不使用代码库中的排序函数来解决这道题吗？
//你能想出一个仅使用常数空间的一趟扫描算法吗？

// 双指针（同时将0和1调整至slice开头）（这个好理解）
func sortColors(nums []int) []int {
	n := len(nums)
	p0, p1 := 0, 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 { // 说明p0最开始的值是1，之前将p0的1交换到i，需要将i交换到p1
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
	return nums
}

// 双指针，将0,2交换到两头
func sortColors2(nums []int) []int {
	n := len(nums)
	p0, p2 := 0, n-1

	for i := 0; i < p2; {
		if nums[i] == 0 && i != p0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		} else if nums[i] == 2 && i != p2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		} else {
			i++
		}
	}
	return nums
}
