package main

import (
	"fmt"
	"runtime"
	"log"
)

func main()  {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	fmt.Println(1)
	where()
	fmt.Println(2)
	where()

}