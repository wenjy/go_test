package main

import (
	"fmt"
)

type Tester interface {
	I1()
}
type a1 struct {
}

func (a *a1) T1() {
	fmt.Println("test")
}

func main() {
	var a a1
	a.T1()
}
