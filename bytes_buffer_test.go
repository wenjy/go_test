package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func main() {

}

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlus("127.0.0.1", "10000", "username", "FZlhufwTZ0s7RYWZ")
	}
}

func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuffer("127.0.0.1", "10000", "username", "FZlhufwTZ0s7RYWZ")
	}
}

func BenchmarkStringFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringFmt("127.0.0.1", "10000", "username", "FZlhufwTZ0s7RYWZ")
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoin("127.0.0.1", "10000", "username", "FZlhufwTZ0s7RYWZ")
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilder("127.0.0.1", "10000", "username", "FZlhufwTZ0s7RYWZ")
	}
}

func StringPlus(ip, port, username, password string) string {
	return ip + port + username + password
}

func StringBuffer(ip, port, username, password string) string {
	var buffer bytes.Buffer
	buffer.WriteString(ip)
	buffer.WriteString(port)
	buffer.WriteString(username)
	buffer.WriteString(password)
	return buffer.String()
}

func StringFmt(ip, port, username, password string) string {
	return fmt.Sprint(ip, port, username, password)
}

func StringJoin(ip, port, username, password string) string {
	s := []string{ip, port, username, password}
	return strings.Join(s, "")
}

func StringBuilder(ip, port, username, password string) string {
	var b strings.Builder
	b.WriteString(ip)
	b.WriteString(port)
	b.WriteString(username)
	b.WriteString(password)
	return b.String()
}
