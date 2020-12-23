package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	go p()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	w()
	time.Sleep(time.Second)
}

func w() {
	ch <- 1
	close(ch)
	ch <- 2
}

func p() {
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		}
	}
}
