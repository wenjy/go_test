package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

type MyFloat float64

// Go 没有类。然而，仍然可以在结构体类型上定义方法。
func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat) Abs1() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex2{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs1())

	v1 := &Vertex2{3, 4}
	v1.Scale(5)
	fmt.Println(v1, v1.Abs())
}
