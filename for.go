package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for ; sum1 < 10; {
		sum1 += sum1
	}
	fmt.Println(sum1)

	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}