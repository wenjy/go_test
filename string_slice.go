package main

import "fmt"

func main() {
	// 修改字符串，必须先将字符串转换成字节数组
	// 然后再通过修改数组中的元素值来达到修改字符串的目的，最后将字节数组转换回字符串格式
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"

	fmt.Println(s2)
}
