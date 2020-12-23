package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8082")

	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024*1024) // 1MB
	count := 0
	num := 10
	start := time.Now().Unix()
	for i := 0; i < num; i++ {
		n, err := conn.Write(buf)
		if err != nil {
			// handle error
			fmt.Println(err)
			return
		}
		count += n
		fmt.Println("count", count, time.Now().Unix()-start)
	}

	for {

	}
}
