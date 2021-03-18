package main

import (
	"fmt"
	"reflect"
)

type conn1 struct {
	fd *fd1
}

type fd1 struct {
	pfd pfd1
}

type pfd1 struct {
	Sysfd int
}

// net.Conn Mock
type MockedConn struct {
	conn conn1
}

func NewMockedConn() *MockedConn {
	return &MockedConn{
		conn: conn1{
			fd: &fd1{
				pfd: pfd1{Sysfd: 1},
			},
		},
	}
}

func main() {
	c := NewMockedConn()
	tcpConn := reflect.Indirect(reflect.ValueOf(c)).FieldByName("conn")
	fmt.Println(tcpConn)
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	fmt.Println(pfdVal.FieldByName("Sysfd").Int())
}
