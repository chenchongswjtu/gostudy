package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var line string
	//fmt.Scanln(&line)
	//line = strings.ReplaceAll(line, ",", "")
	//var max int
	//for i := 0; i < len(line); i++ {
	//	if line[i] == '0' {
	//		pre := strings.Index(line, "1")
	//	}
	//}

	//fmt.Println(strings.Index("10000100101", "1"))
	//fmt.Println(strings.LastIndex("10000100101", "1"))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)
}
