package main

// 暴力
func leastInterval(tasks []byte, n int) int {
	countAll := len(tasks)
	counts := make(map[byte]int)
	nexts := make(map[byte]int)

	for _, task := range tasks {
		counts[task]++
	}

	for k := range counts {
		nexts[k] = -1 // 初始化每个key下一个最小的位置
	}

	var result []byte
	var index int

	for countAll > 0 {
		isSet := false
		for k, count := range counts {
			if countAll > 0 && count > 0 && index > nexts[k] {
				result = append(result, k)
				countAll--
				counts[k]--
				nexts[k] = nexts[k] + n + 1 // 重新计算位置
				index++
				isSet = true
				break
			}
		}

		if !isSet { // 没有放置key的放#
			result = append(result, '#')
			index++
		}

	}
	//fmt.Println(result)
	return len(result)
}
