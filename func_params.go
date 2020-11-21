package main

import "fmt"

func main() {
	var strss = []string{
		"qwr",
		"234",
		"yui",
		"cvbc",
	}
	test1(strss...) //切片被打散传入
}

func test1(args ...string) { //可以接受任意个string参数
	for _, v := range args {
		fmt.Println(v)
	}
}
