package main

import (
	"fmt"
	"net"
)

func main() {
	hostport := "www.baidu.com"
	h, p, err := net.SplitHostPort(hostport)
	fmt.Println(h, p, err)

	hostport = "www.baidu.com:443"
	h, p, err = net.SplitHostPort(hostport)
	fmt.Println(h, p, err)
}
