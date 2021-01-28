package main

import (
	"fmt"
	"time"
)

type minuteSpeed [60]uint32

func main() {
	ms := new(minuteSpeed)
	now := time.Now().Unix()

	s1 := now % 60
	s2 := uint8(now % 60)
	fmt.Println(s1, s2, time.Now().Second())

	mu := now - s1
	fmt.Println(mu)

	ms[s1] = 1
	ms[s2] += 1
	fmt.Println(ms)

	t4 := now % 3600 / 60
	t5 := time.Now().Minute()
	fmt.Println(t4, t5)
}
