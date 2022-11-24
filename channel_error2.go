package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
可能不止有一个发送者（或者不只发送一次），所以这个方案要求，缓存队列的容量需要和发送次数一致。
一旦缓存队列容量被用完后，再有发送者发送就会阻塞发送者 goroutine。如果恰好此时接收者退出了，那么仍然至少会有一个 goroutine 无法退出，从而造成内存泄漏

在代码 (1) 处，接收者的 goroutine 中，我们加了一个判断语句。
这会让代码 (2) 处的 channel 还没被接收完就退出了接收者 goroutine。尽管代码 (3) 处有缓存，但是因为发送 channel 在 for 循环中，缓存队列很快就会被占满，阻塞在第 101 的位置。
所以这种情况我们要使用一个额外的 stop channel 来终结发送者所在的 goroutine

不论发送者发送一次还是多次，如果接收者所在 goroutine 能够在接收完 channel 中的数据之后结束，那么就不会造成内存泄漏；或者说接收者能够在发送者停止发送后再结束，就不会造成内存泄露

如果接收者需要在 channel 关闭之前提前退出，为防止内存泄漏，在发送者与接收者发送次数是一对一时，应设置 channel 缓冲队列为 1；
在发送者与接收者的发送次数是多对多时，应使用专门的 stop channel 通知发送者关闭相应 channel
*/
func main() {
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	chanLeakOfMemory2()
	// chanLeakOfMemory22()
	time.Sleep(time.Second * 3) // 等待 goroutine 执行，防止过早输出结果
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
}

func chanLeakOfMemory2() {
	ich := make(chan int, 100) // (3)
	// sender
	go func() {
		defer close(ich)
		for i := 0; i < 10000; i++ {
			ich <- i
			time.Sleep(time.Millisecond) // 控制一下，别发太快
		}
	}()
	// receiver
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		for i := range ich { // (2)
			if ctx.Err() != nil { // (1)
				fmt.Println(ctx.Err())
				return
			}
			fmt.Println(i)
		}
	}()
}

func chanLeakOfMemory22() {
	ich := make(chan int, 100)
	stopCh := make(chan struct{})
	// sender
	go func() {
		defer close(ich)
		for i := 0; i < 10000; i++ {
			select {
			case <-stopCh:
				return
			case ich <- i:
			}
			time.Sleep(time.Millisecond) // 控制一下，别发太快
		}
	}()
	// receiver
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		for i := range ich {
			if ctx.Err() != nil {
				fmt.Println(ctx.Err())
				close(stopCh)
				return
			}
			fmt.Println(i)
		}
	}()
}
