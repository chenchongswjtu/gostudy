package main

import (
	"fmt"
	"math/rand"
)

// 最大堆，优先队列
type maxPQ struct {
	data []int // [0]不使用, 数据从1开始
	size int   // 数据数量
}

func newMaxPQ() *maxPQ {
	return &maxPQ{
		data: []int{0},
		size: 0,
	}
}

func (pq *maxPQ) getSize() int {
	return pq.size
}

func (pq *maxPQ) isEmpty() bool {
	return pq.size == 0
}

func (pq *maxPQ) insert(value int) {
	pq.data = append(pq.data[:pq.size+1], value)
	pq.size++

	pq.down2up()
}

func (pq *maxPQ) down2up() {
	var i = pq.size
	for i > 1 && pq.data[i] > pq.data[i/2] {
		pq.data[i], pq.data[i/2] = pq.data[i/2], pq.data[i]
		i = i / 2
	}
}

func (pq *maxPQ) delMax() int {
	if pq.isEmpty() {
		panic("max pq is empty, can not del")
	}
	e := pq.data[1]

	pq.data[1] = pq.data[pq.size]
	pq.size--

	if pq.size < len(pq.data)/2 {
		pq.resize()
	}

	pq.up2down()

	return e
}

func (pq *maxPQ) getMax() int {
	if pq.size > 0 {
		return pq.data[1]
	}

	panic("empty can not getMax")
}

func (pq *maxPQ) up2down() {
	var n int
	var i = 1
	for 2*i <= pq.size {
		if 2*i+1 <= pq.size && pq.data[2*i] < pq.data[2*i+1] {
			n = 2*i + 1
		} else {
			n = 2 * i
		}

		if pq.data[i] > pq.data[n] {
			break
		}

		pq.data[i], pq.data[n] = pq.data[n], pq.data[i]

		i = n
	}
}

func (pq *maxPQ) print() {
	fmt.Printf("size is %d\n", pq.getSize())
	for i := 1; i <= pq.size; i++ {
		fmt.Printf("i = %d, v = %d\n", i, pq.data[i])
	}
}

func (pq *maxPQ) resize() {
	var data1 = make([]int, pq.size+1)
	data1[0] = 0
	for i := 1; i <= pq.size; i++ {
		data1[i] = pq.data[i]
	}
	pq.data = nil
	pq.data = data1
}

func lowN(a []int, n int) []int {
	if len(a) <= n {
		return a
	}

	pq := newMaxPQ()
	var res []int
	for i := 0; i < len(a); i++ {
		if pq.size < n {
			pq.insert(a[i])
		} else {
			if a[i] < pq.getMax() {
				pq.delMax()
				pq.insert(a[i])
			}
		}
	}

	for i := 1; i <= pq.size; i++ {
		res = append(res, pq.data[i])
	}

	return res
}

func genRand(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(1000))
	}

	return res
}

func main() {
	//pq := newMaxPQ()
	//pq.insert(3)
	//pq.print()
	//pq.insert(4)
	//pq.print()
	//pq.insert(5)
	//pq.print()
	//pq.insert(6)
	//pq.print()
	//pq.insert(1)
	//pq.print()
	//pq.insert(2)
	//pq.print()
	//fmt.Println(pq.delMax())
	//fmt.Println("####", len(pq.data), cap(pq.data))
	//fmt.Println(pq.delMax())
	//fmt.Println("####", len(pq.data), cap(pq.data))
	//fmt.Println(pq.delMax())
	//fmt.Println("####", len(pq.data), cap(pq.data))
	//fmt.Println(pq.delMax())
	//fmt.Println("####", len(pq.data), cap(pq.data))
	//fmt.Println(pq.delMax())
	//fmt.Println("####", len(pq.data), cap(pq.data))
	//fmt.Println(pq.delMax())
	//fmt.Println("####", len(pq.data), cap(pq.data))
	//pq.print()

	var a = genRand(10)
	fmt.Println(a)
	fmt.Println(lowN(a, 4))
}
