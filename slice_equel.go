package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	s1 := []byte("abcd")
	s2 := []byte("abcd")
	s3 := []byte("abcde")

	fmt.Println("s1 == s2 ?", reflect.DeepEqual(s1, s2))
	fmt.Println("s1 == s3 ?", reflect.DeepEqual(s1, s3))

	// 转换为string来比较
	fmt.Println("s1 == s2 ?", bytes.Equal(s1, s2))
	fmt.Println("s1 == s3 ?", bytes.Equal(s1, s3))
}
