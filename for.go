package main

import (
	"fmt"
)

var forMap = map[string]int{
	"a": 1,
	"b": 2,
}

func main() {
	//var a = []byte{1, 2, 3, 4, 5, 6}
	//fmt.Println(a[0:5])
	fmt.Println(len([]byte{}))
	var splitLen = 16352
	var dataLen = 16352*3 + 1

	var end = 0

	stop := false
	for start := 0; start < dataLen; start += splitLen {
		if stop {
			break
		}
		end += splitLen
		if end > dataLen {
			end = dataLen
			stop = true
		}

		fmt.Println(start, end)
	}

	return
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 < 10 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// like while
	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// key val
	for key, val := range forMap {
		fmt.Printf("key:%s val:%d\n", key, val)
	}

	// key
	for key := range forMap {
		fmt.Printf("key:%s\n", key)
	}

	// val
	for _, val := range forMap {
		fmt.Printf("val:%d\n", val)
	}

	// for string
	forStr := "我是开发者"
	for pos, char := range forStr {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	forArr := []int{1, 2, 3, 4}

	// Reverse a
	for i, j := 0, len(forArr)-1; i < j; i, j = i+1, j-1 {
		forArr[i], forArr[j] = forArr[j], forArr[i]
	}
	fmt.Println(forArr)

}
