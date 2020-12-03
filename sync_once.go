package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

// 保证once只执行一次，无论你是否更换once.Do(xx)这里的方法,这个sync.Once块只会执行一次。
// 整个程序，只会执行onces()方法一次,onced()方法是不会被执行的
func main() {

	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(1 * 1e9)
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
