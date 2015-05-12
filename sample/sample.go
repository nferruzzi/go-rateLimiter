package main

import (
	"log"
	"time"

	"github.com/nferruzzi/go-rateLimiter"
)

func main() {
	limiter := RateLimiter.NewRateLimiter(time.Second / 10)

	n := time.Now()
	for i := 0; i < 10; i++ {
		log.Print("Scheduled:", i)

		e := i
		f := func() {
			log.Printf("%d: %v", e, time.Now().Sub(n))
		}

		limiter.Go(f)
	}

	time.Sleep(time.Second)
}
