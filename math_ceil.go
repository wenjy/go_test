package main

import (
	"fmt"
	"math"
)

func main() {
	chunk := 129 / 128
	fmt.Println(chunk)
	if chunk == 0 {
		chunk = 1
	}
	a := int(math.Ceil(float64(chunk)))
	fmt.Println(a)
}
