package main

import (
	"fmt"
	"time"
)

func main() {
    ch1 := make(chan int)
    go pump(ch1)       // pump hangs
	// fmt.Println(<-ch1) // prints only 0 没有再接收了
	go suck(ch1)
	time.Sleep(1e9)
}

func pump(ch chan int) {
    for i := 0; ; i++ {
        ch <- i
    }
}

func suck(ch chan int) {
    for {
        fmt.Println(<-ch)
    }
}
