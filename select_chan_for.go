package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool, 1)

	go func() {
		timeCh := time.After(3 * time.Second)
		select {
		case <-timeCh:
			fmt.Println("timed after")
			ch1 <- 1
			close(ch1)
			// ch2 <- true
		}
	}()

	go func() {
		for {
			select {
			case i, ok := <-ch1:
				if !ok {
					fmt.Println("before send ch2")
					// ch2不是异步（没有容量），需要go协程才行，不然是阻塞的
					//go func() {
					ch2 <- true
					//}()
					fmt.Println("after send ch2")
					break
				}
				fmt.Println(i)
			case <-ch2:
				fmt.Println("stop")
				return
			}
		}
	}()

	time.Sleep(time.Second * 10)
}
