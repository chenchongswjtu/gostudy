package main

import "fmt"

func main() {
Loop:
	for j := 0; j < 3; j++ {
		fmt.Println(j)
		for a := 0; a < 5; a++ {
			fmt.Println(a)
			if a > 3 {
				continue Loop
			}
		}
	}
}
