package main

import "fmt"

type Aa struct {
	Data int
	buf  []byte
}

func main() {
	a := &Aa{Data: 1}
	if a.buf == nil {
		fmt.Println(fmt.Sprintf("111%p", a))
	}
	fmt.Println(fmt.Sprintf("%p", a))
	f1(a)
	f2(a)
}

func f1(a *Aa) {
	fmt.Println(fmt.Sprintf("%p", a))
}

func f2(a *Aa) {
	fmt.Println(fmt.Sprintf("%p", a))
}
