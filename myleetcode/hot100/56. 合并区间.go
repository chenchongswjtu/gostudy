package main

import "sort"

//56. 合并区间
//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
//
//
//示例 1：
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2：
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//提示：
//
//1 <= intervals.length <= 104
//intervals[i].length == 2
//0 <= starti <= endi <= 104

func merge1(intervals [][]int) [][]int {
	// 先按照start排序，从小到大
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		// 如果前一个end小于后一个start，什么都不做
		if intervals[i-1][1] < intervals[i][0] {
			continue
		}

		// 如果前一个end大于等于后一个end，合并，i--
		if intervals[i-1][1] >= intervals[i][1] {
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
		} else {
			// 如果前一个end小于后一个end，合并，i--
			tmp := []int{intervals[i-1][0], intervals[i][1]}
			intervals = append(append(intervals[:i-1], tmp), intervals[i+1:]...)
			i--
		}

	}

	return intervals
}
