package main

import (
	"fmt"
	"math"
)

//var n, t, c, s float64 = 9581, 1, 100, 0.99

func main() {
	//s := 0.99
	//cs := []float64{200.0, 400.0, 800.0, 1000.0}
	//for _, c := range cs {
	//	for i := 1; i <= 3; i++ {
	//		t := float64(i)
	//		fmt.Printf("N=%f, T=%f, C=%f, FN=%f\n", N(t, c, s), t, c, s)
	//	}
	//}

	n, t, c := 10000.0, 3.7, 100.0

	fmt.Println(S(n, t, c))
	fmt.Println(S1(n, t, c))

	//fmt.Println(T(10000.0, 100.0, 0.963))
}

// 参数n 账户数
// 参数t 交易间隔时间(单位秒)
// 参数c 并发数
// 参数s 交易成功率
// s=(1-1/n)**(c*t-1)

// 计算账户数n
func N(t, c, s float64) float64 {
	var n = 1.0
	for {
		fn := math.Pow(1.0-1.0/n, c*t-1)
		if fn >= s {
			return n
		} else if fn >= s {

		} else {
			n = n + 1
		}
	}
}

// 计算交易成功率s
func S(n, t, c float64) float64 {
	return math.Pow(1.0-1.0/n, c*t-1.0)
}

// 计算c
func C(n, t, s float64) float64 {
	l := math.Log(s) / math.Log(1.0-1.0/n)
	return (l + 1.0) / t
}

// 计算t
func T(n, c, s float64) float64 {
	l := math.Log(s) / math.Log(1.0-1.0/n)
	return (l + 1.0) / c
}

func S1(N, T, C float64) float64 {
	Namda0 := math.Pow(1.0-1.0/N, C)
	Namda1 := 1.0 / N * math.Pow(1.0-1.0/N, C-1)
	FN := math.Pow(Namda0, T) + T*Namda1*math.Pow(Namda0, T-1)
	return FN
}
