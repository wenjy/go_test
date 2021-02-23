package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 与互斥量不同，条件变量的作用并不是保证在同一时刻仅有一个线程访问某一个共享数据，而是在对应的共享数据的状态发生变化时，通知其他因此而被阻塞的线程。
	// 条件变量总是与互斥量组合使用。互斥量为共享数据的访问提供互斥支持，而条件变量可以就共享数据的状态的变化向相关线程发出通知。
	cond := sync.NewCond(new(sync.Mutex))
	condition := 0

	// Consumer
	go func() {
		for {
			cond.L.Lock()
			for condition == 0 {
				cond.Wait()
			}
			condition--
			fmt.Printf("Consumer: %d\n", condition)
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// Producer
	for {
		time.Sleep(time.Second)
		cond.L.Lock()
		for condition == 3 {
			cond.Wait()
		}
		condition++
		fmt.Printf("Producer: %d\n", condition)
		cond.Signal()
		cond.L.Unlock()
	}
}
