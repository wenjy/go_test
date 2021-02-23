package main

import "fmt"

func main() {
	var test = make([]int, 10)

	var j int
	for i := 0; i < 5; i++ {
		test[j] = i
		j++
	}

	fmt.Println(test)
	fmt.Println(test[:j])
	for _, num := range test[:j] {
		fmt.Println(num)
	}
}
