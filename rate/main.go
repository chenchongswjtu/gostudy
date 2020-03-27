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

	// 1秒間に実行する回数
	ntimes := 5

	n := rate.Every(time.Second / time.Duration(ntimes))
	limiter := rate.NewLimiter(n, 1)

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
