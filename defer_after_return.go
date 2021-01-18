package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("before")
	}()

	return

	defer func() {
		fmt.Println("after")
	}()
}
