package main

import "fmt"

// x = x0 - f(x0)/f `(x0)       ( f `(x) 为函数 f(x)  的一阶导数 f `(x) != 0)

func main() {
	fmt.Println(getCube(8))
	fmt.Println(getSqrt(9))
}

// f(x) = x*x*x-n
func getCube(n float64) float64 {
	var x = 1.0
	var x1 = x - (x*x*x-n)/(3*x*x)
	for x-x1 > 0.0000001 || x-x1 < -0.0000001 {
		x = x1
		x1 = x - (x*x*x-n)/(3*x*x)
	}
	return x1
}

// f(x) = x*x-n
func getSqrt(n float64) float64 {
	var x = 1.0
	var x1 = x - (x*x-n)/(2*x)
	for x-x1 > 0.0000001 || x-x1 < -0.0000001 {
		x = x1
		x1 = x - (x*x-n)/(2*x)
	}
	return x1
}
