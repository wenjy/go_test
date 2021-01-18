package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	stopDone := make(chan struct{})
	stopWaitGroup := sync.WaitGroup{}
	stopWaitGroup.Add(3)

	go func() {
		defer close(stopDone)
		stopWaitGroup.Wait()
	}()

	go func() {
		select {
		case <-stopDone:
			return
		case <-time.After(time.Second * 10):
			fmt.Println("stop timeout")
			os.Exit(0)
		}
	}()

	done1()
	stopWaitGroup.Done()
	done2()
	stopWaitGroup.Done()
	done3()
	stopWaitGroup.Done()

	fmt.Println("stoped")
	os.Exit(0)
}

func done1() {
	time.Sleep(time.Second * 3)
	fmt.Println("done1")
}

func done2() {
	time.Sleep(time.Second * 2)
	fmt.Println("done2")
}

func done3() {
	time.Sleep(time.Second * 3)
	fmt.Println("done3")
}
