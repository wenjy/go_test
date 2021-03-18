package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go func() {
		timeCh := time.After(3 * time.Second)
		select {
		case <-timeCh:
			fmt.Println("timed after")
			//ch1 <- 1
			close(ch1)
			close(ch2)
			// ch2 <- true
		}
	}()

	go func() {
		for {
			select {
			case i, ok := <-ch1:
				if !ok {
					fmt.Println("stop ch1")
					break
				}
				fmt.Println(i)
			case <-ch2:
				fmt.Println("stop ch2")
				return
			}
		}
	}()

	time.Sleep(time.Second * 10)
}
