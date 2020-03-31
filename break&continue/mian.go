package main

import "fmt"

// break到Loop，之后不继续执行多层循环，Loop必须在for之前
func breakTest() {
Loop:
	for j := 0; j < 3; j++ {
		fmt.Println(j)
		for a := 0; a < 5; a++ {
			fmt.Println(a)
			if a > 3 {
				break Loop
			}
		}
	}
}

// continue到Loop，之后继续执行多层循环，Loop必须在for之前
func continueTest() {
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

func main() {
	breakTest()

	fmt.Println("##############")

	continueTest()
	//output:
	//0
	//0
	//1
	//2
	//3
	//4
	//##############
	//0
	//0
	//1
	//2
	//3
	//4
	//1
	//0
	//1
	//2
	//3
	//4
	//2
	//0
	//1
	//2
	//3
	//4
}
