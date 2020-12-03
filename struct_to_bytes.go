package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	ID   int32
	Data int32
	//Data2 string
}

func main() {
	var buf = []byte{1, 2, 3, 4, 1, 2, 3, 4, "a"}

	var ptestStruct *myStruct = *(**myStruct)(unsafe.Pointer(&buf))

	fmt.Println(unsafe.Sizeof(myStruct{}))
	fmt.Println(ptestStruct.ID, ptestStruct.Data) // 67305985 6
}
