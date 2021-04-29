package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	//"github.com/petermattis/goid"
)

func main() {
	fmt.Println(goID())
}

//goroutine 1 [running]:
//main.goID(0x569ee0)
//D:/gopath/src/github

func goID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	//fmt.Println(string(b))
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return n
}
