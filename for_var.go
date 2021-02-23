package main

import "fmt"

func main() {
	var retry int

	for retry = 0; retry < 3; retry++ {
		break
		fmt.Println(retry)
	}

	fmt.Println(retry)
}
