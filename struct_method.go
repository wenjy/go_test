package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())                 // 22
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20)) // 42

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem()) // 7

	fmt.Println(IntVector{1, 2, 3}.Sum()) // 6
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
