package main

import (
	"fmt"
	"unsafe"
)

type UnsafeS struct {
	A uint32
	B uint64
}

func main() {
	bytes1 := []byte{104, 101, 108, 108, 111}

	p := unsafe.Pointer(&bytes1) //强制转换成unsafe.Pointer，编译器不会报错
	str := *(*string)(p)         //然后强制转换成string类型的指针，再将这个指针的值当做string类型取出来
	fmt.Println(str)             //输出 "hello"

	fmt.Println(unsafe.Offsetof(UnsafeS{}.B)) // 8
	fmt.Println(unsafe.Sizeof(UnsafeS{}.A))   // 4
	fmt.Println(unsafe.Sizeof(UnsafeS{}))     // 16
	fmt.Println(unsafe.Alignof(UnsafeS{}))    // 8 对齐大小
}
