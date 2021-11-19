package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// 转换8进制utf-8字符串到中文
// eg: `\346\200\241` -> 怡
func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}

func main() {
	s1 := "\350\264\246\346\210\267\344\275\231\351\242\235\344\270\215\350\266\263" // 字面量
	s2 := `\350\264\246\346\210\267\344\275\231\351\242\235\344\270\215\350\266\263` // 原始字符串

	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)

	// 转化 s2
	s3 := convertOctonaryUtf8(s2)
	fmt.Println("s3 =", s3)
}
