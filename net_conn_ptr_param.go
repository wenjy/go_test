package main

import (
	"fmt"
	"net"
	"time"
)

type MyConn struct {
}

func (mc MyConn) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (mc MyConn) Write(b []byte) (n int, err error) {
	return 0, nil
}

func (mc MyConn) Close() error {
	return nil
}

func (mc MyConn) LocalAddr() net.Addr {
	return nil
}

func (mc MyConn) RemoteAddr() net.Addr {
	return nil
}

func (mc MyConn) SetDeadline(t time.Time) error {
	return nil
}

func (mc MyConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (mc MyConn) SetWriteDeadline(t time.Time) error {
	return nil
}

type MyConn1 struct {
}

func (mc *MyConn1) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (mc *MyConn1) Write(b []byte) (n int, err error) {
	return 0, nil
}

func (mc *MyConn1) Close() error {
	return nil
}

func (mc *MyConn1) LocalAddr() net.Addr {
	return nil
}

func (mc *MyConn1) RemoteAddr() net.Addr {
	return nil
}

func (mc *MyConn1) SetDeadline(t time.Time) error {
	return nil
}

func (mc *MyConn1) SetReadDeadline(t time.Time) error {
	return nil
}

func (mc *MyConn1) SetWriteDeadline(t time.Time) error {
	return nil
}

func main() {
	var conn net.Conn
	f(conn)

	var myConn MyConn
	f(myConn)

	myConn1 := new(MyConn1)
	f(myConn1)
}

func f(c net.Conn) {
	fmt.Println(c)
}
