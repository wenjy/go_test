package main

import (
	"fmt"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

func main() {
	date := "2021-12-30 11:11:11"

	time1, err := time.Parse(timeLayout, date)

	fmt.Println(err)
	fmt.Println(time1.Unix())
}
