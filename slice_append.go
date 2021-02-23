package main

import "fmt"

func main() {
	var test []int

	fmt.Println("cap", cap(test)) // 0
	fmt.Println("len", len(test)) // 0

	test = append(test, 1)
	fmt.Println("cap", cap(test)) // 1
	fmt.Println("len", len(test)) // 1

	test = append(test, 2)
	fmt.Println("cap", cap(test)) // 2
	fmt.Println("len", len(test)) // 2

	test = append(test, 3)
	fmt.Println("cap", cap(test)) // 4
	fmt.Println("len", len(test)) // 3

	var test2 = make([]int, 2)
	fmt.Println("cap2", cap(test2)) // 2
	fmt.Println("len2", len(test2)) // 2

	test2 = append(test2, 3)
	fmt.Println("cap2", cap(test2)) // 4
	fmt.Println("len2", len(test2)) // 3

	test2 = append(test2, 4)
	fmt.Println("cap2", cap(test2)) // 4
	fmt.Println("len2", len(test2)) // 4

	test2 = append(test2, 5)
	fmt.Println("cap2", cap(test2)) // 8
	fmt.Println("len2", len(test2)) // 5
}
