package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "ab:cd"
	res := strings.Split(str1, ":")

	fmt.Println(res)
}
