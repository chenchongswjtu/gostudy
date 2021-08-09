package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	ctx1, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	slow(ctx1)
	fmt.Println(time.Since(t1))
	time.Sleep(3 * time.Second)
	fmt.Println(time.Since(t1))
	fmt.Println("main return")
}

func slow(ctx context.Context) {
	var i = 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println(i)
			return
			//i++
			//time.Sleep(1 * time.Second)
		}
	}
}
