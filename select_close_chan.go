package main

import (
	"fmt"
	"time"
)

type Ab struct {
	data int
}

var ch = make(chan *Ab, 5)

// 关闭通道，select 读取的是通道默认值
func main() {
	go p()

	ch <- &Ab{data: 1}
	ch <- &Ab{data: 2}
	ch <- &Ab{data: 3}

	close(ch)
	time.Sleep(time.Microsecond * 1000)
}

func p() {
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		}
	}
}
