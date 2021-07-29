package main

import (
	"context"
	"log"
	"os"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lmicroseconds)

	// 1秒之内执行的次数
	ntimes := 5
	limiter := rate.NewLimiter(rate.Limit(ntimes), 1)

	ctx := context.Background()

	log.Println("--- Start ---")
	for i := 0; i < 20; i++ {
		if err := limiter.Wait(ctx); err != nil {
			log.Fatalln(err)
		}

		log.Printf("Do work %02d", i+1)

		if 10 <= (i + 1) {
			time.Sleep(200 * time.Millisecond)
		}

		if limiter.Allow() {
			log.Println("Allow() true")
		} else {
			log.Println("Allow() false")
		}
	}
}
