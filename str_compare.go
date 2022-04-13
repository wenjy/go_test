package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "147"
	str2 := "148"
	str3 := "147"
	fmt.Println(strings.Compare(str1, str2)) // < -1
	fmt.Println(strings.Compare(str2, str1)) // > 1
	fmt.Println(strings.Compare(str1, str3)) // = 0
}
