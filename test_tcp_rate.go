package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

func main() {
	limiter = rate.NewLimiter(1024*128, 1024*128)
	ln, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	var count int
	buf := make([]byte, 65536)
	start := time.Now().Unix()
	for {
		parent := context.Background()
		ctx, cancelFunc := context.WithTimeout(parent, time.Second)
		err := limiter.WaitN(ctx, 65536)
		if err != nil {
			cancelFunc()
			fmt.Println("cannel", err)
			continue
		}
		rn, err := c.Read(buf)
		count += rn
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("count", count, time.Now().Unix()-start)
	}
}
