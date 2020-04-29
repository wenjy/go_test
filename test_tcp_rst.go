package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var host = flag.String("host", "172.16.1.49", "host")
var port = flag.String("port", "9999", "port")

func main()  {
	conn, err := net.Dial("tcp", *host+":"+*port)
	fmt.Println(*host+":"+*port)
	if err != nil {
		fmt.Println("err1")
		os.Exit(1)
	}

	defer conn.Close()

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		fmt.Println("err2")
		os.Exit(1)
	}
	tcpConn.SetLinger(0)
}