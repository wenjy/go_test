package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

/*
这是 go channel 导致内存泄漏的经典场景。根据输出结果（开始有两个 goroutine，结束时有三个 goroutine），我们可以知道，直到测试函数结束前，仍有一个 goroutine 没有退出。
原因是由于 (1) 处创建的 errCh 是不含缓存队列的 channel，如果 channel 只有发送方发送，那么发送方会阻塞；如果 channel 只有接收方，那么接收方会阻塞。
我们可以看到由于没有发送方往 errCh 发送数据，所以 (4) 处代码一直阻塞。直到 (3) 处超时后，打印“超时”，函数退出，(4) 处代码都未接收成功。
而 (2) 处的所在的 goroutine 在“超时”被打印后，才开始发送。由于外部的 goroutine 已经退出了，errCh 没有接收者，导致 (2) 处一直阻塞。
因此 (2) 处代码所在的协程一直未退出，造成了内存泄漏。如果代码中有许多类似的代码，或在 for 循环中使用了上述形式的代码，随着时间的增长会造成多个未退出的 gorouting，最终导致程序 OOM


我们只需要为 channel 增加一个缓存队列。即把 (1) 处代码改为 errCh := make(chan error, 1) 即可

可能会有人想要使用 defer close(errCh) 关闭 channel

由于 (2) 处代码没有接收者，所以一直阻塞。直到 close(errCh) 运行，(2) 处仍在阻塞。这导致关闭 channel 时，仍有 goroutine 在向 errCh 发送。然而在 golang 中，在向 channel 发送时不能关闭 channel，否则会 panic。因此这种方式是错误的

又或在 (5) 处 goroutine 的第一句加上 defer close(errCh)。由于 (2) 处阻塞， defer close(errCh) 会一直得不到执行。因此也是错误的。 即便对调 (2) 处和 (4) 处的发送者和接收者，也会因为 channel 关闭，导致输出无意义的零值

*/

func main() {
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	chanLeakOfMemory()
	time.Sleep(time.Second * 3) // 等待 goroutine 执行，防止过早输出结果
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
}

func chanLeakOfMemory() {
	errCh := make(chan error) // (1)
	// errCh := make(chan error, 1) // 正确
	go func() { // (5)
		time.Sleep(2 * time.Second)
		errCh <- errors.New("chan error") // (2)
		fmt.Println("finish sending")
	}()

	var err error
	select {
	case <-time.After(time.Second): // (3) 大家也经常在这里使用 <-ctx.Done()
		fmt.Println("超时")
	case err = <-errCh: // (4)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(nil)
		}
	}
}
