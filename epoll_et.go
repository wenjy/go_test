package main

import (
	"fmt"
	"net"
	"reflect"
	"strings"
	"syscall"
)

const EPOLLET = 1 << 31

type Epoll struct {
	fd          int
	events      []syscall.EpollEvent
	waitNum     int //每次等待的事件数量
	waitTimeout int //等待超时ms
}

func NewEpoll(waitNum, waitTimeout int) (*Epoll, error) {
	fd, err := syscall.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &Epoll{
		fd:          fd,
		events:      make([]syscall.EpollEvent, waitNum),
		waitNum:     waitNum,
		waitTimeout: waitTimeout,
	}, nil
}

// 添加读事件
// fd 需要添加到epoll的文件描述符
// enableET 是否启用et模式
func (e *Epoll) Add(fd int, enableET bool) error {
	var ev *syscall.EpollEvent
	if enableET {
		ev = &syscall.EpollEvent{Events: syscall.EPOLLIN | syscall.EPOLLHUP | EPOLLET, Fd: int32(fd)}
	} else {
		ev = &syscall.EpollEvent{Events: syscall.EPOLLIN | syscall.EPOLLHUP, Fd: int32(fd)}
	}
	err := syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, ev)
	if err != nil {
		return err
	}

	return nil
}

// 删除读事件
func (e *Epoll) Remove(fd int) error {
	err := syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}

	return nil
}

// 等待可读的连接
func (e *Epoll) Wait(wf func(fd int)) error {
retry:
	n, err := syscall.EpollWait(e.fd, e.events, e.waitTimeout)
	if err != nil {
		if err == syscall.EINTR {
			goto retry
		}
		return err
	}

	for i := 0; i < n; i++ {
		wf(int(e.events[i].Fd))
	}

	return nil
}

// 关闭epoll
func (e *Epoll) Close() {
	syscall.Close(e.fd)
}

var epoller *Epoll

func main() {
	var err error
	epoller, err = NewEpoll(100, 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}

	go startEpoll()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
		fdVal := tcpConn.FieldByName("fd")
		pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
		fd := pfdVal.FieldByName("Sysfd").Int()

		if err := epoller.Add(int(fd), true); err != nil {
			fmt.Println("device epoller add error: " + err.Error())
			return
		}
	}
}

func startEpoll() {
	for {
		err := epoller.Wait(func(fd int) {
			fmt.Println(fd)
		})
		if err != nil {
			// epoll 已关闭
			if strings.Contains(err.Error(), "bad file descriptor") {
				return
			}
			fmt.Println("epoller.Wait err:" + err.Error())
			continue
		}
	}
}
