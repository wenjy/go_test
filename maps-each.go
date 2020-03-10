package main

import "fmt"

func main() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	scene["ip%5B0%5D"] = 123
	for k, v := range scene {
	    fmt.Println(k, v)
	}

	for k := range scene {
	    fmt.Println(k)
	}
}
