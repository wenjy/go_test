package main

import (
	"sync"
	"fmt"
)

type Info struct {
    mu sync.Mutex
    name string
}

func main() {
	info := Info{name:"bbb"}
	fmt.Println(info.name)
	Update(&info)
	fmt.Println(info.name)
}

func Update(info *Info) {
    info.mu.Lock()
    // critical section:
    info.name = "aaa"// new value
    // end critical section
    info.mu.Unlock()
}