package main

import (
	"math/rand"
	"time"
	"unsafe"
	"fmt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytes(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// 防止产生相同的随机数
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random 随机生成指定长度字符串
func Random(n int) string {
	//rand.Seed(time.Now().UnixNano())
	return randStringBytes(n)
}

func main() {
	go func() {
		fmt.Println(Random(32))
	}()
	go func() {
		fmt.Println(Random(32))
	}()
	go func() {
		fmt.Println(Random(32))
	}()
	
	fmt.Println(Random(32))
	fmt.Println(Random(32))
	fmt.Println(Random(32))
	fmt.Println(Random(32))
	fmt.Println(Random(32))

	time.Sleep(time.Second)
}
