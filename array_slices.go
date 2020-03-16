package main

import "fmt"

func main() {
	// 切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度以及切片容量
    var arr1 [6]int
    var slice1 []int = arr1[2:5] // 从索引2开始，但不包括索引为5的

    // load the array with integers: 0,1,2,3,4,5
    for i := 0; i < len(arr1); i++ {
        arr1[i] = i
    }

    // print the slice
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    }

    fmt.Printf("The length of arr1 is %d\n", len(arr1))
    fmt.Printf("The length of slice1 is %d\n", len(slice1))
    fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

    // grow the slice，相比slice1往后移了一位
    slice1 = slice1[0:4]
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    }
    fmt.Printf("The length of slice1 is %d\n", len(slice1))
    fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

    // grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
	
	b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	b1 := b[1:4] // 0 l a
	for i := 0; i < len(b1); i++ {
        fmt.Println(i, b1[i])
	}
	
	b2 := b[:2] // g o
	for i := 0; i < len(b2); i++ {
        fmt.Println(i, b2[i])
	}

	b3 := b[2:] // l a n g
	for i := 0; i < len(b3); i++ {
        fmt.Println(i, b3[i])
	}

	b4 := b[:] // g o l a n g
	for i := 0; i < len(b4); i++ {
        fmt.Println(i, b4[i])
	}
}