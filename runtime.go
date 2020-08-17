package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	fmt.Println(1)
	where()
	fmt.Println(2)
	where()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
