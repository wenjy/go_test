package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

const DATA_BUFSIZE = 4096

func main() {

	fd, _ := windows.Socket(windows.AF_INET, windows.SOCK_STREAM, windows.IPPROTO_TCP)

	addr := windows.SockaddrInet4{
		Port: 9100,
		Addr: [4]byte{172, 16, 1, 49},
	}

	connectErr := windows.Connect(fd, &addr)

	defer windows.Close(fd)

	if connectErr != nil {
		fmt.Println("connectErr", connectErr)
		return
	}

	buf := []byte("test")

	wsaBuf := windows.WSABuf{
		Len: uint32(len(buf)),
		Buf: &buf[0],
	}

	var flag, flag1 uint32 = 0, 0
	var sent, recvd uint32 = 0, 0
	//var overlapped windows.Overlapped
	//var croutine byte

	sendErr := windows.WSASend(fd, &wsaBuf, 1, &sent, flag, nil, nil)
	if sendErr != nil {
		fmt.Println(sendErr)
		return
	}

	fmt.Println("send", sent)

	recvBuf := [DATA_BUFSIZE]byte{}
	wsaBuf1 := windows.WSABuf{
		Len: uint32(DATA_BUFSIZE),
		Buf: &recvBuf[0],
	}

	recvErr := windows.WSARecv(fd, &wsaBuf1, 1, &recvd, &flag1, nil, nil)
	if recvErr != nil {
		fmt.Println("recvErr", recvErr)
		return
	}

	fmt.Println("recvd", recvd)
	if recvd > 0 {
		fmt.Println("wsaBuf1", string(recvBuf[:recvd]))
	}
}
