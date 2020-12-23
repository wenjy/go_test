package main

import "fmt"

func main() {
	buf := []byte("qwertyuio")
	bufLen := len(buf)
	start := 0
	n := 3
	fmt.Println(bufLen)
	for {
		if bufLen < n {
			n = bufLen
		}
		//bufLen -= n
		fmt.Println(buf[start : start+n])
		start += n
		bufLen -= n
		if bufLen == 0 {
			break
		}
	}
	fmt.Println(start)
}
