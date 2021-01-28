package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type foo struct {
	a uint32
}

var fooMap = make(map[int]*foo, 10)

func main() {
	fooMap[1] = &foo{}
	fooMap[2] = &foo{}

	for i := 0; i <= 2000; i++ {
		// go incr(1)
		// go incr(2)
		go atomicIncr(1)
		go atomicIncr(2)
	}

	time.Sleep(time.Second * 5)
	fmt.Println(fooMap[1])
	fmt.Println(fooMap[2])
}

func incr(id int) {
	fooMap[id].a += 1
}

func atomicIncr(id int) {
	atomic.AddUint32(&fooMap[id].a, uint32(1))
}
