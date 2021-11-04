package main

import "fmt"

const (
	AAA int = iota
	BBB
	CCC
)

const (
	DDD = 1 << iota
	EEE
	FFF
)

func main() {
	fmt.Println(AAA)
	fmt.Println(BBB)
	fmt.Println(CCC)
	fmt.Println(DDD)
	fmt.Println(EEE)
	fmt.Println(FFF)
}
