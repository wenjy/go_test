package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func main() {

	fd, _ := windows.Socket(windows.AF_INET, windows.SOCK_DGRAM, windows.IPPROTO_UDP)

	addr := windows.SockaddrInet4{
		Port: 9200,
		Addr: [4]byte{172, 16, 1, 49},
	}

	buf := []byte("test")

	if err := windows.Sendto(fd, buf, 0, &addr); err != nil {
		fmt.Println("SendtoErr", err)
	}
	return
}
