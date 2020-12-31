package main

import (
	"fmt"
	"runtime"
	"time"
)

var c chan int

//func handle(int) {}

func main() {
	go func() {
		timeCh := time.After(3 * time.Second)
		fmt.Println("NumGoroutine", runtime.NumGoroutine())
		select {
		case <-c:

		case <-timeCh:
			fmt.Println("timed out")
		}
	}()
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	time.Sleep(time.Second * 4)
}
