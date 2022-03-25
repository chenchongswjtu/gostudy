package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(a73("xyxyXX"))
	fmt.Println(a73("abababb"))
}

type pair struct {
	key   uint8
	count int
}

func a73(str string) string {
	var m = make(map[uint8]int)
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}

	var pairs = make([]pair, len(m))
	var index int
	for k, v := range m {
		pairs[index] = pair{key: k, count: v}
		index++
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].key < pairs[j].key
		} else {
			return pairs[i].count < pairs[j].count
		}
	})

	var lastUpper int
	for i := 0; i < len(pairs); i++ {
		if pairs[i].key >= 'A' && pairs[i].key <= 'Z' {
			lastUpper = i
		}
	}

	pairs = append(pairs[lastUpper+1:], pairs[:lastUpper+1]...)

	var ret string
	for _, p := range pairs {
		ret += fmt.Sprintf("%s:%d;", string(p.key), p.count)
	}
	return ret
}
