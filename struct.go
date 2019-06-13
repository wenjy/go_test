package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Vertex1 struct {
	X, Y int
}

var (
	v1 = Vertex1{1, 2}  // 类型为 Vertex
	v2 = Vertex1{X: 1}  // Y:0 被省略
	v3 = Vertex1{}      // X:0 和 Y:0
	a  = &Vertex1{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, a, a.X, v2, v3)
}
