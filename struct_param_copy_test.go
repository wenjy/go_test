package main

import "testing"

type AuthKey string
type InfoKey string
type Test22 struct {
	f1 string
	f6 int
	f2 int
	f3 int64
	f4 AuthKey
	f5 InfoKey
}

func BenchmarkTest1(b *testing.B) {
	as := Test22{
		"qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		12345,
		12345,
		12345,
		AuthKey("127.0.0.112345test1qqqqqqqqqqqqqqqq"),
		InfoKey("123456:127.0.0.1:12345"),
	}
	for i := 0; i < b.N; i++ {
		func21(as)
	}
}

func BenchmarkTest2(b *testing.B) {
	as := &Test22{
		"qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		12345,
		12345,
		12345,
		AuthKey("127.0.0.112345test1qqqqqqqqqqqqqqqq"),
		InfoKey("123456:127.0.0.1:12345"),
	}
	for i := 0; i < b.N; i++ {
		func22(as)
	}
}

func func21(as Test22)  {}
func func22(as *Test22) {}

func main() {}
