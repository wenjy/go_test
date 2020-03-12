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

	// 函数可以将其他函数调用作为它的参数，
	// 只要这个被调用函数的返回值个数、返回值类型和返回值的顺序与调用函数所需求的实参是一致的
	fmt.Println(f1(f2()))
}

func f1(a, b int) int {
	return a + b
}

func f2() (a, b int) {
	return 1,2
}
