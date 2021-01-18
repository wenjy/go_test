package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringBytesMaskImpr(n int) string {
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

var num = 100000
var cache = make(map[string]bool, num)
var mutex = sync.Mutex{}

var runNum uint32

func main() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(4)

	var i int
	for i = 0; i < num; i++ {
		go addCache()
	}

	time.Sleep(time.Second * 1)
	fmt.Println(runNum)
}

func addCache() {
	str := RandStringBytesMaskImpr(32)

	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := cache[str]; ok {
		fmt.Println(str)
		return
	}
	cache[str] = true
	atomic.AddUint32(&runNum, 1)
}
