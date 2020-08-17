package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

func main() {
	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)

	addr := syscall.SockaddrInet4{
		Port: 30001,
		Addr: [4]byte{172, 16, 1, 49},
	}

	connectErr := syscall.Connect(fd, &addr)

	defer syscall.Close(fd)

	if connectErr != nil {
		fmt.Println(connectErr)
		return
	}

	res, err := syscall.Write(fd, []byte("test"))

	/* file, err := MakeTunFile(uintptr(fd))
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := file.WriteString("test") */

	fmt.Println(res) // 0
	fmt.Println(err) // The parameter is incorrect.
}

// MakeTunFile returns an os.File object from a TUN file descriptor `fd`.
func MakeTunFile(fd uintptr) (*os.File, error) {
	if fd < 0 {
		return nil, errors.New("TUN file descriptor error")
	}
	file := os.NewFile(fd, "")
	if file == nil {
		return nil, errors.New("open TUN file descriptor error")
	}
	return file, nil
}
