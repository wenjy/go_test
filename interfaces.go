package main

import (
	"fmt"
	"math"
	"os"
)

type Abser interface {
	Abs() float64
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var a Abser
	f := MyFloat1(-math.Sqrt2)
	v := Vertex1{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())

	var w Writer

	// os.Stdout 实现了 Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}

type MyFloat1 float64

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex1 struct {
	X, Y float64
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
