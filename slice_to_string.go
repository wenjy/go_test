package main

import "fmt"

type myString string

func main() {
	s1 := []byte("abcd")

	s2 := myString(s1)

	fmt.Println(s2) // abcd
	s1[0] = 101

	fmt.Println(s2)         // abcd
	fmt.Println(string(s1)) // ebcd
}
