package main

import "fmt"

type Vertex4 struct {
	X int
	Y int
}

type Vertex3 struct {
	X, Y int
}

var (
	v1 = Vertex3{1, 2}  // 类型为 Vertex
	v2 = Vertex3{X: 1}  // Y:0 被省略
	v3 = Vertex3{}      // X:0 和 Y:0
	a  = &Vertex3{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(Vertex4{1, 2})

	v := Vertex4{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, a, a.X, v2, v3)
}
