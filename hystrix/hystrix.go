package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	hystrix.Go("get_baidu",
		func() error { // 熔断器操作
			// talk to other services
			_, err := http.Get("https://www.baidu.com/")
			if err != nil {
				fmt.Println("get error", err)
				return err
			}
			return nil
		},
		func(err error) error { // 熔断器失败之后的操作
			fmt.Println("get an error, handle it", err)
			return nil
		})

	time.Sleep(10 * time.Second) // 调用Go方法就是起了一个goroutine，这里要sleep一下，不然看不到效果
}
