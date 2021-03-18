package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU()
	fmt.Println(num)
	old := runtime.GOMAXPROCS(num)
	fmt.Println(old)
}
