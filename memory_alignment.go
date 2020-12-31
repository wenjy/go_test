package main

import (
	"fmt"
	"unsafe"
)

// S的是8字节对齐
type S struct {
	A uint32   // 4byte 需对齐8byte
	B uint64   // 8byte
	C uint64   // 8byte
	D uint64   // 8byte
	E struct{} // 0byte 结构体尾部size为0的变量(字段)会被分配内存空间进行填充，原因是如果不给它分配内存，该变量(字段)指针将指向一个非法的内存空间(类似C/C++的野指针)
}

func main() {
	fmt.Println(unsafe.Alignof(S{}))    // 8 对齐大小
	fmt.Println(unsafe.Offsetof(S{}.E)) // 32 偏移位置
	fmt.Println(unsafe.Sizeof(S{}.E))   // 0 内存大小
	fmt.Println(unsafe.Sizeof(S{}))     // 40
}
