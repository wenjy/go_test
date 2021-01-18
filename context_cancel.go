package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	l, err := NewTCPListener(ctx, net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9500}, handle)

	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		var ch chan int
		select {
		case <-time.After(time.Second * 10):
			if l != nil {
				fmt.Println("cancel")
				l.Close()
				cancel()
				return
			}
		case <-ch:
		}
	}()

	for {
		time.Sleep(time.Second * 1)
	}
}

type TcpHandle func(ctx context.Context, c net.Conn)

// NewTCPListener 创建新的TCP Listen
func NewTCPListener(ctx context.Context, addr net.TCPAddr, process TcpHandle) (listener net.Listener, err error) {
	listener, err = net.Listen("tcp", addr.String())
	if err != nil {
		return nil, err
	}

	AcceptTcp(ctx, listener, process)

	return listener, nil
}

// AcceptTcp 接受请求
func AcceptTcp(ctx context.Context, l net.Listener, process TcpHandle) {

	select {
	case <-ctx.Done():
		fmt.Println("Done1")
		return
	default:
	}
	for {
		c, err := l.Accept()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			if strings.Contains(err.Error(), "the mux has closed") {
				break
			}
			fmt.Println("Accept err: " + err.Error())
			continue
		}

		go process(ctx, c)
	}
}

func handle(ctx context.Context, c net.Conn) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done2")
			c.Close()
			return
		default:
			buf := make([]byte, 1024)
			rLen, err := c.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			c.Write(buf[:rLen])
		}
	}
}
