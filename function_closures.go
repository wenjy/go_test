package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func(i int) int {
	return func(i int) int {
		return i
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
