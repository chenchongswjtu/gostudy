package main

import (
	"fmt"
	"strconv"
)

func main() {
	count("HMg2(H2ON3)3N2")
}

// HMg2(H2ON3)3N2
func count(s string) map[string]int {
	if len(s) == 0 {
		return nil
	}

	i := 0
	n := len(s)
	for {
		if i >= n {
			break
		}
		if i > 0 && ('A' <= s[i] && s[i] <= 'Z') && ('A' <= s[i-1] && s[i-1] <= 'Z') {
			s = s[:i] + "1" + s[i:]
			i = i + 2
		}
		i++
		n = len(s)
	}

	fmt.Println(s)

	var starts, ends []int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			starts = append(starts, i)
		}

		if s[i] == ')' {
			ends = append(ends, i)
		}
	}

	fmt.Println(starts)
	fmt.Println(ends)

	k := make(map[string]int)
	for i := 0; i < len(starts); i++ {
		t := s[starts[i]+1 : ends[i]]
		j := ends[i] + 1
		for {
			if '0' <= s[j] && s[j] <= '9' {
				j++
			} else {
				break
			}
		}

		n, _ := strconv.Atoi(s[ends[i]+1 : j])

		k[t] = n
	}

	t := s[0:starts[0]]
	k[t] = 1
	t = s[ends[len(ends)-1]+2:]
	k[t] = 1
	fmt.Println(k)

	res := make(map[string]int)
	for kk, v := range k {
		pre := 0
		for i := 0; i < len(kk); i++ {
			if '0' <= kk[i] && kk[i] <= '9' {
				n, _ := strconv.Atoi(kk[i : i+1])
				res[kk[pre:i]] += n * v
				pre = i + 1
			}
		}
	}

	fmt.Println(res)

	return res
}
