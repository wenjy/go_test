package main

import "fmt"

func main() {
	b := make([]byte, 10)
	fmt.Println(b)
	b1 := b[:3]
	b2 := b[3:6]
	//b3 := b[6:]
	b3 := make([]byte, 10)

	copy(b1, []byte("aaa"))
	copy(b2, []byte("bbb"))
	copy(b3, b2)
	b = b[3:6]
	//copy(b, b[3:6])
	//copy1(b3)
	fmt.Println(b)
	fmt.Println(b3)
}

func copy1(b []byte) {
	copy(b, []byte("cccc"))
}
