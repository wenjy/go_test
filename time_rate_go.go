package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(2, 10)

func main() {
	n := 6
	fmt.Println(limiter.Burst())
	for i := 0; i < 100; i++ {
		go waitNFun1(n)
		go waitNFun2(n)
		go waitNFun3(n)
	}

	time.Sleep(time.Second * 30)
}

func waitNFun1(n int) {
	limiter.WaitN(context.Background(), n)
	fmt.Println(limiter.Burst())
	fmt.Println("func1")
}

func waitNFun2(n int) {
	limiter.WaitN(context.Background(), n)
	fmt.Println("func2")
}

func waitNFun3(n int) {
	limiter.WaitN(context.Background(), n)
	fmt.Println("func3")
}
