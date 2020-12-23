package main

import (
	"context"
	"fmt"
	
	"time"
	"golang.org/x/time/rate"
)

func main() {
	ch := make(chan int)
	limiter := rate.NewLimiter(1, 10)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
			start := time.Now().Unix()
			err := limiter.WaitN(context.Background(), 10)
			if err != nil {
				fmt.Println("err", err)
				break
			}
			println(time.Now().Unix() - start)
		case <-ch:
			println("case2")
		}
	}
}