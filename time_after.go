package main

import (
	"fmt"
	"time"
)

var c chan int

//func handle(int) {}

func main() {
	select {
	case <-c:

	case <-time.After(3 * time.Second):
		fmt.Println("timed out")
	}
}
