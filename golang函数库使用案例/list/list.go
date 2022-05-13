package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 声明链表
	l := list.New()
	// 数据添加到尾部
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)

	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	// 遍历
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}

	l.PushFront(0)
	l1 := list.New()
	l1.PushBack(34)
	l1.PushBack(35)
	l.PushBackList(l1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
