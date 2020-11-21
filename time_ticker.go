package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)
	var i int
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t, "hello world")
			if i >= 5 {
				ticker.Stop()
				return
			}
			i++
		}
	}
}
