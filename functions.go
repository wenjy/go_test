package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var php, c bool
var d, f int = 1, 2

func main() {
	var i int
	e := "aaa"
	fmt.Println(add(2, 5))
	fmt.Println(add2(3, 5))
	a, b := swap("world", "hello")
	fmt.Printf("%s %s\n", a, b)
	fmt.Println(split(11))
	fmt.Println(php, c, i, d, f, e)
}
