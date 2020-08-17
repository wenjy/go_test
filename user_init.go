package main

import (
	"fmt"

	"./trans"
)

func main() {
	twoPi := trans.Pi * 2
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}
