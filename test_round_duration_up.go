package main

import (
	"fmt"
	"time"
)

func main() {
	t := roundDurationUp(time.Second*10, time.Second)
	fmt.Println(t)
}

func roundDurationUp(d time.Duration, to time.Duration) time.Duration {
	return (d + to - 1) / to
}
