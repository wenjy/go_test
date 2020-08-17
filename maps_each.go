package main

import "fmt"

func main() {
	// map 无序
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	scene["ip%5B0%5D"] = 123
	for k, v := range scene {
		fmt.Println(k, v)
	}

	// 只获取key
	for k := range scene {
		fmt.Println(k)
	}

	// 只获取value
	for _, value := range scene {
		fmt.Println(value)
	}
}
