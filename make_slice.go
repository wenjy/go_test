package main

import "fmt"

func main() {
	// make 的使用方式是：func make([]T, len, cap)，其中 cap 是可选参数
	var slice1 []int = make([]int, 3, 5)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
